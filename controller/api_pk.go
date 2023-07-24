package controller

import (
	"CqepcAuto/api"
	"CqepcAuto/api/cqepc"
	"CqepcAuto/api/dingtalk"
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// ApiPkAssist 补评课程
func ApiPkAssist(ctx *gin.Context) {
	// 获取参数
	classroomId := ctx.PostForm("classroomId")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验参数
	if err := utils.SafeApp.SafeVerify(gin.H{
		"classroomId": classroomId,
		"timestamp":   timestamp,
		"sign":        sign,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	// 补评课程
	// 1. 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取数据库用户信息错误，Error: " + err.Error(),
			"data": "",
		})
		return
	}

	//2. 获取课程信息
	classInfo, err := model.TodayNew(model.Today{Uid: userInfo.Uid, PId: strings.TrimSpace(classroomId)}).GetByUidAndPID()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取数据库用户信息错误，Error: " + err.Error(),
			"data": "",
		})
		return
	}
	var tempClassData global.ResTodayClassData
	json.Unmarshal([]byte(classInfo.Data), &tempClassData)
	if time.Now().Unix() < utils.ToolApp.DateToUnix(tempClassData.NodeBegin) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  fmt.Sprintf("课程还未开始，请课程开始后( %v )再次尝试", tempClassData.NodeBegin),
			"data": "",
		})
		return
	}
	if classInfo.IsPk == "success" || classInfo.IsPk == "danger" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "课程已评价，无需再次评价",
			"data": "",
		})
		return
	}

	// 3. 获取评课内容
	pkMsg, err := api.ApiGroupApp.CqepcGroup.ClassPK.New(cqepc.ClassPK{
		Token:       userInfo.PkToken,
		ClassroomId: strings.TrimSpace(classroomId),
	}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取评课信息失败，Error: " + err.Error(),
			"data": "",
		})
		return
	}
	if pkMsg.Code != 200 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "网络请求获取评课信息失败，Error: " + pkMsg.Msg,
			"data": "",
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
		ClassroomId:       strings.TrimSpace(classroomId),
		CoincidenceDegree: pkRandMsg,
	}).Send()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "评课失败，Error: " + err.Error(),
			"data": "",
		})
		return
	}
	if res.Code != 200 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "网络请求评课失败，Error: " + pkMsg.Msg,
			"data": "",
		})
		return
	}

	err = model.TodayNew(model.Today{
		Uid:       classInfo.Uid,
		PId:       classInfo.PId,
		Data:      classInfo.Data,
		PkData:    pkMsg.Msg,
		IsPk:      "danger",
		CreatedAt: classInfo.CreatedAt,
	}).UpdateByUidAndPid()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "修改课程状态失败，请再次重试! Error: " + pkMsg.Msg,
			"data": "",
		})
		return
	}

	// 发送钉钉消息
	api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
		Token:  userInfo.NToken,
		Secret: userInfo.NSecret,
		Phone:  []string{userInfo.NPhone},
	}).SendMarkdown("补评课程通知", []string{
		"# 补评课程成功",
		"------",
		fmt.Sprintf("您补评课程 **%v(%v)** 请求已发送成功，该功能比较危险，请谨慎使用！\n", tempClassData.CourseName, tempClassData.CourseCode),
		"------",
		"## 评课内容",
		"- 上课是否能听懂: 是",
		"- 老师是否上课: 是",
		"- 是否有作业: 是",
		"- 备注: ",
		fmt.Sprintf("- 吻合度: %v", pkRandMsg),
	})

	// 写入日志
	model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     userInfo.Uid,
		Msg:     "补评课程成功",
		MsgType: "pk_assist",
		Data:    fmt.Sprintf("课程ID： %v, 上课是否能听懂: %v, 老师是否上课: %v, 是否有作业: %v, 备注： , 吻合度: %v", tempClassData.Id, "是", "是", "是", pkRandMsg),
	})
	model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     userInfo.Uid,
		Msg:     "补评课程成功",
		MsgType: "ding_talk_msg",
		Data:    fmt.Sprintf("课程ID： %v, 上课是否能听懂: %v, 老师是否上课: %v, 是否有作业: %v, 备注： , 吻合度: %v", tempClassData.Id, "是", "是", "是", pkRandMsg),
	})

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}
