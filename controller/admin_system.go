package controller

import (
	"CqepcAuto/api"
	"CqepcAuto/api/cqepc"
	"CqepcAuto/api/dingtalk"
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AdminSystem(ctx *gin.Context) {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.String(http.StatusOK, "获取数据库信息异常，请再次重试！")
		return
	}

	userInfo.NToken = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v", userInfo.NToken)

	// 返回数据
	ctx.HTML(http.StatusOK, "admin/system.html", gin.H{
		"user": userInfo,
		"time": time.Now().Unix(),
	})
}

func AdminSystemPost(ctx *gin.Context) {
	// 获取参数
	studentID := ctx.PostForm("StudentID")
	studentPassword := ctx.PostForm("StudentPassword")
	studentOpenid := ctx.PostForm("StudentOpenid")
	dingTalkWebhook := ctx.PostForm("DingTalkWebhook")
	dingTalkSecret := ctx.PostForm("DingTalkSecret")
	dingTalkPhone := ctx.PostForm("DingTalkPhone")
	sysLoginType := ctx.PostForm("SysLoginType")
	sysCode := ctx.PostForm("SysCode")
	sign := ctx.PostForm("sign")
	timestamp := ctx.PostForm("timestamp")

	// 校验参数
	err := installIndexValidate(ctx, studentID, studentPassword, studentOpenid, dingTalkWebhook, dingTalkSecret, sysLoginType, sysCode, sign, timestamp, dingTalkPhone)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	if studentOpenid == "NULL" {
		studentOpenid = ""
	}

	// 用户登录cqepc
	var userInfo global.ResLogin
	if strings.TrimSpace(sysLoginType) == "openid" {
		userInfo, err = api.ApiGroupApp.CqepcGroup.Login.New(cqepc.Login{
			Openid: strings.TrimSpace(studentOpenid),
		}).AutoLogin()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "网络请求错误，Error: " + err.Error(),
				"data": "",
			})
			return
		}
	} else if strings.TrimSpace(sysLoginType) == "account" {
		userInfo, err = api.ApiGroupApp.CqepcGroup.Login.New(cqepc.Login{
			UserName: strings.TrimSpace(studentID),
			PassWord: strings.TrimSpace(studentPassword),
			Openid:   strings.TrimSpace(studentOpenid),
		}).UPLogin()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "网络请求错误，Error: " + err.Error(),
				"data": "",
			})
			return
		}
	}

	if userInfo.Code != 200 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "评课信息配置: " + userInfo.Msg,
			"data": "",
		})
		return
	}

	// 获取数据库信息
	dbInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取数据库信息失败，请再次重试！",
			"data": "",
		})
		return
	}

	// 更新数据库
	index := strings.Index(strings.TrimSpace(dingTalkWebhook), "access_token") + len("access_token") + 1
	token := dingTalkWebhook[index:]
	err = model.UserNew(model.User{
		Uid:        dbInfo.Uid,
		SId:        userInfo.Data.Username,
		SPass:      strings.TrimSpace(studentPassword),
		SOpenid:    strings.TrimSpace(studentOpenid),
		SName:      userInfo.Data.Nickname,
		SClass:     userInfo.Data.ClassName,
		SClassCode: userInfo.Data.ClassCode,
		PkToken:    userInfo.Data.Token,
		LoginType:  strings.TrimSpace(sysLoginType),
		NToken:     token,
		NSecret:    dingTalkSecret,
		NPhone:     dingTalkPhone,
		IsStop:     "register",
	}).UpdatedByUid()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "写入信息失败，请再次重试！",
			"data": "",
		})
		return
	}

	// 发送消息通知用户
	api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
		Token:  token,
		Secret: strings.TrimSpace(dingTalkSecret),
		Phone:  []string{strings.TrimSpace(dingTalkPhone)},
	}).SendMarkdown("系统配置更新通知", []string{
		"# 系统配置更新通知",
		"----",
		fmt.Sprintf("您在 %v 更新了系统配置\n - 学号: %v\n - 姓名: %v\n - 班级: %v\n", time.Now().Format("2006-01-02 15:04:05"), userInfo.Data.Username, userInfo.Data.Nickname, userInfo.Data.ClassName),
		"",
		"---",
		"**查看更多内容，请前往 Cqepc Auto 后台管理系统查看详情**",
	})

	// 写入日志
	model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     dbInfo.Uid,
		Msg:     "系统配置更新通知",
		MsgType: "ding_talk_msg",
		Data:    fmt.Sprintf("您在 %v 更新了系统配置\n - 学号: %v\n - 姓名: %v\n - 班级: %v\n", time.Now().Format("2006-01-02 15:04:05"), strings.TrimSpace(studentID), userInfo.Data.Nickname, userInfo.Data.ClassName),
	}).Add()

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}
