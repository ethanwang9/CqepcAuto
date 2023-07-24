package core

import (
	"CqepcAuto/api"
	"CqepcAuto/api/cqepc"
	"CqepcAuto/api/dingtalk"
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

// name: 定时任务
// author: axel wong
// desc: 定时任务计划

type CronDo struct{}

// CronInit 初始化定时器
func CronInit() *cron.Cron {
	return cron.New(cron.WithSeconds())
}

// CronDoNew 初始化定时任务
func CronDoNew() *CronDo {
	return &CronDo{}
}

// DO 定时任务集合
func (d *CronDo) DO() *CronDo {
	// 获取每日课表
	global.APP_CRON.AddFunc("27 0 7 * * ?", d.GetEveryDayClassTable)

	// 登录
	global.APP_CRON.AddFunc("0 0 7-22 * * ?", d.GetLogin)

	// 更新评课数据
	global.APP_CRON.AddFunc("17 0 22 * * ?", d.UpdatePkData)

	// 自动评课
	global.APP_CRON.AddFunc("23 0 8-22 * * ?", d.SendPK)

	// BUG修复
	global.APP_CRON.AddFunc("2 0 0 * * ?", d.GetLogin)
	global.APP_CRON.AddFunc("30 0 0 * * ?", d.GetEveryDayClassTable)

	return d
}

// Run 运行定时任务
func (d *CronDo) Run() {
	global.APP_CRON.Start()
}

// GetEveryDayClassTable 获取每日课表
func (d *CronDo) GetEveryDayClassTable() {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		global.APP_LOG.Error("定时任务@获取每日课表-获取数据库信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@获取每日课表-获取数据库信息失败")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-获取数据库信息失败",
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 删除今日冗余信息
	err = model.TodayNew(model.Today{Uid: userInfo.Uid}).DelTodayClassByUid()
	if err != nil {
		global.APP_LOG.Error("定时任务@获取每日课表-删除今日冗余信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText(fmt.Sprintf("定时任务@获取每日课表-删除今日冗余信息\n\nError: %v", err.Error()))

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-删除今日冗余信息",
			MsgType: "ding_talk_msg",
			Data:    fmt.Sprintf("定时任务@获取每日课表-删除今日冗余信息\n\nError: %v", err.Error()),
		})

		return
	}

	// 获取课表
	today, err := api.ApiGroupApp.CqepcGroup.ClassToday.New(cqepc.ClassToday{Token: userInfo.PkToken}).Get()
	if err != nil {
		global.APP_LOG.Error("定时任务@获取每日课表-获取今日课表失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText(fmt.Sprintf("定时任务@获取每日课表-获取今日课表失败\n\nError: %v", err.Error()))

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-获取今日课表失败",
			MsgType: "ding_talk_msg",
			Data:    fmt.Sprintf("定时任务@获取每日课表-获取今日课表失败\n\nError: %v", err.Error()),
		})

		return
	}

	if today.Code != 200 {
		global.APP_LOG.Error("定时任务@获取每日课表-获取今日课表请求失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText(fmt.Sprintf("定时任务@获取每日课表-获取今日课表请求失败\n\nError: %v", today.Msg))

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-获取今日课表请求失败",
			MsgType: "ding_talk_msg",
			Data:    fmt.Sprintf("定时任务@获取每日课表-获取今日课表请求失败\n\nError: %v", today.Msg),
		})

		return
	}

	// 写入数据库
	for _, v := range today.Data {
		tempData, _ := json.Marshal(v)
		err := model.TodayNew(model.Today{
			Uid:  userInfo.Uid,
			PId:  v.Id,
			Data: string(tempData),
			IsPk: "no",
		}).Add()
		if err != nil {
			global.APP_LOG.Error("定时任务@获取每日课表-写入今日课表到数据库失败", zap.Error(err), zap.Any("pk", v))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText(fmt.Sprintf("定时任务@获取每日课表-写入今日课表到数据库失败\n\n"+
				" - 课程名: %v\n\n"+
				" - 错误信息: %v",
				v.ClassName,
				err.Error(),
			))

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取每日课表-写入今日课表到数据库失败",
				MsgType: "ding_talk_msg",
				Data:    fmt.Sprintf("定时任务@获取每日课表-写入今日课表到数据库失败\n\n"+" - 课程名: %v\n\n"+" - 错误信息: %v", v.ClassName, err.Error()),
			})
		}
	}
}

// GetLogin 登录
func (d *CronDo) GetLogin() {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		global.APP_LOG.Error("定时任务@获取登录状态-获取数据库信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@获取登录状态-获取数据库信息失败")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取登录状态-获取数据库信息失败",
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 获取登录状态
	var loginData global.ResLogin
	if userInfo.LoginType == "openid" {
		loginData, err = api.ApiGroupApp.CqepcGroup.Login.New(cqepc.Login{
			Openid: userInfo.SOpenid,
		}).AutoLogin()
		if err != nil {
			global.APP_LOG.Error("定时任务@获取登录状态-自动登录失败！", zap.Error(err))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText("定时任务@获取登录状态-自动登录失败")

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取登录状态-自动登录失败",
				MsgType: "ding_talk_msg",
			})

			return
		}
		if loginData.Code != 200 {
			global.APP_LOG.Error("定时任务@获取登录状态-自动登录失败！", zap.String("error", loginData.Msg))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText("定时任务@获取登录状态-自动登录失败\n\n" +
				"消息：" + loginData.Msg)

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取登录状态-自动登录失败",
				MsgType: "ding_talk_msg",
				Data:    "定时任务@获取登录状态-自动登录失败\n\n" + "消息：" + loginData.Msg,
			})

			return
		}
	} else if userInfo.LoginType == "account" {
		loginData, err = api.ApiGroupApp.CqepcGroup.Login.New(cqepc.Login{
			UserName: userInfo.SId,
			PassWord: userInfo.SPass,
		}).UPLogin()
		if err != nil {
			global.APP_LOG.Error("定时任务@获取登录状态-账号密码登录失败！", zap.Error(err))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText("定时任务@获取登录状态-账号密码登录失败")

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取登录状态-账号密码登录失败",
				MsgType: "ding_talk_msg",
			})

			return
		}
		if loginData.Code != 200 {
			global.APP_LOG.Error("定时任务@获取登录状态-账号密码登录失败！", zap.String("error", loginData.Msg))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText("定时任务@获取登录状态-账号密码登录失败\n\n" +
				"消息：" + loginData.Msg)

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取登录状态-账号密码登录失败",
				MsgType: "ding_talk_msg",
				Data:    "定时任务@获取登录状态-账号密码登录失败\n\n" + "消息：" + loginData.Msg,
			})

			return
		}
	} else {
		// 错误配置
		global.APP_LOG.Error("定时任务@获取登录状态-用户登录类型未无效配置，建议重新安装系统！", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@获取登录状态-用户登录类型未无效配置，建议重新安装系统！")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取登录状态-用户登录类型未无效配置，建议重新安装系统！",
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 写入数据库
	userInfo.PkToken = loginData.Data.Token
	err = model.UserNew(userInfo).UpdatedByUid()
	if err != nil {
		// 错误配置
		global.APP_LOG.Error("定时任务@获取登录状态-写入数据库失败！", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@获取登录状态-写入数据库失败！")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取登录状态-写入数据库失败！",
			MsgType: "ding_talk_msg",
		})

		return
	}
}

// UpdatePkData 更新评课数据
func (d *CronDo) UpdatePkData() {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		global.APP_LOG.Error("定时任务@更新评课数据-获取数据库信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@更新评课数据-获取数据库信息失败")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@更新评课数据-获取数据库信息失败",
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 获取本期课程
	list, err := api.ApiGroupApp.CqepcGroup.ClassList.New(cqepc.ClassList{Token: userInfo.PkToken}).Get()
	if err != nil {
		global.APP_LOG.Error("定时任务@更新评课数据-获取本期课程失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@更新评课数据-获取本期课程失败 Error: " + err.Error())

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@更新评课数据-获取本期课程失败",
			MsgType: "ding_talk_msg",
			Data:    err.Error(),
		})

		return
	}
	if list.Code != 200 {
		global.APP_LOG.Error("定时任务@更新评课数据-获取本期课程请求失败", zap.String("msg", list.Msg))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText(fmt.Sprintf("定时任务@更新评课数据-获取本期课程请求失败\n\nError: %v", list.Msg))

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-获取本期课程请求失败",
			MsgType: "ding_talk_msg",
			Data:    fmt.Sprintf("定时任务@更新评课数据-获取本期课程请求失败\n\nError: %v", list.Msg),
		})

		return
	}

	// 获取评课数据
	pk := make([]global.PkTj, 0)
	for _, v := range list.Data {
		tempData, err := api.ApiGroupApp.CqepcGroup.ClassDetails.New(cqepc.ClassDetails{
			Token:      userInfo.PkToken,
			CourseCode: v.CourseCode,
		}).Get()
		if err != nil {
			global.APP_LOG.Error("定时任务@更新评课数据-获取评课数据失败", zap.Error(err))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText(fmt.Sprintf("定时任务@更新评课数据-获取评课数据失败\n\nError: %v", err.Error()))

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取每日课表-获取今日课表请求失败",
				MsgType: "ding_talk_msg",
				Data:    fmt.Sprintf("定时任务@更新评课数据-获取本期课程请求失败\n\nError: %v", err.Error()),
			})

			return
		}
		if tempData.Code != 200 {
			global.APP_LOG.Error("定时任务@更新评课数据-获取评课数据请求失败", zap.String("data", tempData.Msg))

			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendText(fmt.Sprintf("定时任务@更新评课数据-获取评课数据请求失败\n\nError: %v", tempData.Msg))

			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "定时任务@获取每日课表-获取今日课表请求失败",
				MsgType: "ding_talk_msg",
				Data:    fmt.Sprintf("定时任务@更新评课数据-获取评课数据请求失败\n\nError: %v", tempData.Msg),
			})

			return
		}

		pk = append(pk, global.PkTj{
			CourseCode: v.CourseCode,
			CourseName: v.CourseName,
			Data: struct {
				TotalClassroom  int     `json:"totalClassroom"`
				StartTime       string  `json:"startTime"`
				EndTime         string  `json:"endTime"`
				Join            float64 `json:"join"`
				TotalEvaluation int     `json:"totalEvaluation"`
				JoinStr         string  `json:"joinStr"`
			}{
				TotalClassroom:  tempData.Data.TotalClassroom,
				StartTime:       tempData.Data.StartTime,
				EndTime:         tempData.Data.EndTime,
				Join:            tempData.Data.Join,
				TotalEvaluation: tempData.Data.TotalEvaluation,
				JoinStr:         tempData.Data.JoinStr,
			},
		})
	}

	pkString, _ := json.Marshal(pk)
	err = model.TjNew(model.Tj{
		Uuid: utils.ToolApp.UUID(),
		Uid:  userInfo.Uid,
		Data: string(pkString),
	}).Add()
	if err != nil {
		global.APP_LOG.Error("定时任务@更新评课数据-保存数据到数据库失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText(fmt.Sprintf("定时任务@更新评课数据-保存数据到数据库失败\n\nError: %v", err.Error()))

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@获取每日课表-保存数据到数据库失败",
			MsgType: "ding_talk_msg",
			Data:    fmt.Sprintf("定时任务@更新评课数据-保存数据到数据库失败\n\nError: %v", err.Error()),
		})
	}

	// end
}

// SendPK 发送评课信息
func (d *CronDo) SendPK() {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		global.APP_LOG.Error("定时任务@发送评课信息-获取数据库信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@发送评课信息-获取数据库信息失败")

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@发送评课信息-获取数据库信息失败",
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 获取课表信息
	today, err := model.TodayNew(model.Today{Uid: userInfo.Uid}).GetByUidToday()
	if err != nil {
		global.APP_LOG.Error("定时任务@发送评课信息-获取数据库信息失败", zap.Error(err))

		api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
			Token:  userInfo.NToken,
			Secret: userInfo.NSecret,
			Phone:  []string{userInfo.NPhone},
		}).SendText("定时任务@发送评课信息-获取数据库信息失败 Error: " + err.Error())

		model.LogNew(model.Log{
			Uuid:    utils.ToolApp.UUID(),
			Uid:     userInfo.Uid,
			Msg:     "定时任务@发送评课信息-获取数据库信息失败 Error: " + err.Error(),
			MsgType: "ding_talk_msg",
		})

		return
	}

	// 开始评课
	for _, v := range today {
		var tempData global.ResTodayClassData
		json.Unmarshal([]byte(v.Data), &tempData)

		now := time.Now().Unix()
		begin := utils.ToolApp.DateToUnix(tempData.NodeBegin)
		end := utils.ToolApp.DateToUnix(tempData.NodeEnd)

		// 开始评课
		if now > begin && now < end && v.IsPk == "no" {
			// 获取评课数据
			pkMsg, err := api.ApiGroupApp.CqepcGroup.ClassPK.New(cqepc.ClassPK{
				Token:       userInfo.PkToken,
				ClassroomId: v.PId,
			}).Get()
			if err != nil {
				global.APP_LOG.Error("定时任务@发送评课信息-获取评课信息失败", zap.Error(err))

				api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
					Token:  userInfo.NToken,
					Secret: userInfo.NSecret,
					Phone:  []string{userInfo.NPhone},
				}).SendText("定时任务@发送评课信息-获取评课信息失败 Error: " + err.Error())

				model.LogNew(model.Log{
					Uuid:    utils.ToolApp.UUID(),
					Uid:     userInfo.Uid,
					Msg:     "定时任务@发送评课信息-获取评课信息失败 Error: " + err.Error(),
					MsgType: "ding_talk_msg",
				})

				return
			}
			if pkMsg.Code != 200 {
				global.APP_LOG.Error("定时任务@发送评课信息-获取评课信息请求失败", zap.String("msg", pkMsg.Msg))

				api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
					Token:  userInfo.NToken,
					Secret: userInfo.NSecret,
					Phone:  []string{userInfo.NPhone},
				}).SendText("定时任务@发送评课信息-获取评课信息请求失败 Error: " + pkMsg.Msg)

				model.LogNew(model.Log{
					Uuid:    utils.ToolApp.UUID(),
					Uid:     userInfo.Uid,
					Msg:     "定时任务@发送评课信息-获取评课信息请求失败 Error: " + pkMsg.Msg,
					MsgType: "ding_talk_msg",
				})

				return
			}

			// 获取随机评课
			rand.Seed(time.Now().UnixNano())
			pkRandNumber := rand.Intn(len(pkMsg.Data))
			pkRandMsg := pkMsg.Data[pkRandNumber].UnitTitle
			if pkRandMsg == "NULL" || pkRandMsg == "undefined" {
				pkRandMsg = "其他"
			}

			// 发送消息
			res, err := api.ApiGroupApp.CqepcGroup.ClassSend.New(cqepc.ClassSend{
				Token:             userInfo.PkToken,
				HasUnderstand:     "1",
				HasTeaching:       "1",
				HasSchoolwork:     "1",
				Memo:              "",
				ClassroomId:       v.PId,
				CoincidenceDegree: pkRandMsg,
			}).Send()
			if err != nil {
				global.APP_LOG.Error("定时任务@发送评课信息-发送评课请求发送错误", zap.Error(err))

				api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
					Token:  userInfo.NToken,
					Secret: userInfo.NSecret,
					Phone:  []string{userInfo.NPhone},
				}).SendText("定时任务@发送评课信息-发送评课请求发送错误 Error: " + err.Error())

				model.LogNew(model.Log{
					Uuid:    utils.ToolApp.UUID(),
					Uid:     userInfo.Uid,
					Msg:     "定时任务@发送评课信息-发送评课请求发送错误 Error: " + err.Error(),
					MsgType: "ding_talk_msg",
				})

				return
			}
			if res.Code != 200 {
				global.APP_LOG.Error("定时任务@发送评课信息-评课请求返回错误", zap.String("msg", res.Msg))

				api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
					Token:  userInfo.NToken,
					Secret: userInfo.NSecret,
					Phone:  []string{userInfo.NPhone},
				}).SendText("定时任务@发送评课信息-评课请求返回错误 Error: " + res.Msg)

				model.LogNew(model.Log{
					Uuid:    utils.ToolApp.UUID(),
					Uid:     userInfo.Uid,
					Msg:     "定时任务@发送评课信息-评课请求返回错误 Error: " + res.Msg,
					MsgType: "ding_talk_msg",
				})

				return
			}

			// 修改课程状态
			err = model.TodayNew(model.Today{
				Uid:       v.Uid,
				PId:       v.PId,
				Data:      v.Data,
				PkData:    pkMsg.Msg,
				IsPk:      "success",
				CreatedAt: v.CreatedAt,
			}).UpdateByUidAndPid()
			if err != nil {
				global.APP_LOG.Error("定时任务@发送评课信息-修改课程状态失败", zap.Error(err))

				api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
					Token:  userInfo.NToken,
					Secret: userInfo.NSecret,
					Phone:  []string{userInfo.NPhone},
				}).SendText("定时任务@发送评课信息-修改课程状态失败 Error: " + err.Error())

				model.LogNew(model.Log{
					Uuid:    utils.ToolApp.UUID(),
					Uid:     userInfo.Uid,
					Msg:     "定时任务@发送评课信息-修改课程状态失败 Error: " + err.Error(),
					MsgType: "ding_talk_msg",
				})

				return
			}

			// 发送钉钉消息
			api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
				Token:  userInfo.NToken,
				Secret: userInfo.NSecret,
				Phone:  []string{userInfo.NPhone},
			}).SendMarkdown("自动评课通知", []string{
				"# 评课成功通知",
				"---",
				fmt.Sprintf("**%v** 同学，您在 **%v - %v** 有一节 **%v** 在 **%v教室** 上课，任课老师 **%v** 正在上课中!",
					userInfo.SName,
					tempData.NodeBegin,
					tempData.NodeEnd,
					tempData.CourseName,
					tempData.ClassClassromm,
					tempData.TeacherName,
				),
				"",
				"---",
				"#### 评价内容",
				"- 上课是否能听懂：是",
				"- 老师是否上课：是",
				"- 是否有作业：是",
				"- 意见：无",
				fmt.Sprintf("- 吻合度：%v", pkRandMsg),
				"---",
				"**查看更多内容，请前往 *Cqepc Auto* 后台管理系统查看详情**",
			})

			// 写入日志
			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "补评课程成功",
				MsgType: "pk_assist",
				Data:    fmt.Sprintf("课程ID： %v, 上课是否能听懂: %v, 老师是否上课: %v, 是否有作业: %v, 备注： 无, 吻合度: %v", tempData.Id, "是", "是", "是", pkRandMsg),
			})
			model.LogNew(model.Log{
				Uuid:    utils.ToolApp.UUID(),
				Uid:     userInfo.Uid,
				Msg:     "补评课程成功",
				MsgType: "ding_talk_msg",
				Data:    fmt.Sprintf("课程ID： %v, 上课是否能听懂: %v, 老师是否上课: %v, 是否有作业: %v, 备注： 无, 吻合度: %v", tempData.Id, "是", "是", "是", pkRandMsg),
			})
		}
	}

	// end
}
