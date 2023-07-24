package controller

import (
	"CqepcAuto/api"
	"CqepcAuto/api/dingtalk"
	"CqepcAuto/core"
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InstallVerify(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "install/verify.html", gin.H{
		"version": global.Version,
	})
}

func InstallVerifyPost(ctx *gin.Context) {
	// 获取参数
	studentName := ctx.PostForm("StudentName")
	studentID := ctx.PostForm("StudentID")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 验证参数
	if err := installVerifyValidate(studentName, studentID, timestamp, sign); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}
	if userInfo.SName != strings.TrimSpace(studentName) || userInfo.SId != strings.TrimSpace(studentID) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "真实姓名或学号错误，请检查后再次提交！",
			"data": "",
		})
		return
	}

	// 是否可以发送消息
	isSend, err := model.LogNew(model.Log{Uid: userInfo.NToken}).IsSendDingTalk()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送钉钉消息失败, Error: " + err.Error(),
			"data": "",
		})
		return
	}
	if !isSend {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送钉钉消息过于频繁，请一分钟后再试！",
			"data": "",
		})
		return
	}

	// 生成消息
	msg := []string{
		"# 系统安装成功通知",
		"---",
		"恭喜您，您的 CqepcAuto 系统已安装成功!",
		"",
		"---",
		"**查看更多内容，请前往 *Cqepc Auto* 后台管理系统查看详情**",
	}

	// 写入日志表
	err = model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     userInfo.Uid,
		Msg:     "系统安装成功通知",
		MsgType: "ding_talk_msg",
		Data:    "系统安装成功通知！恭喜您，您的 CqepcAuto 系统已安装成功!",
	}).Add()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "处理数据失败，请再次重试！",
			"data": "",
		})
		return
	}

	// 发送消息通知
	err = api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
		Token:  userInfo.NToken,
		Secret: userInfo.NSecret,
		Phone:  []string{strings.TrimSpace(userInfo.NPhone)},
	}).SendMarkdown("系统安装成功通知", msg)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "钉钉发送消息失败，Error: " + err.Error(),
			"data": "",
		})
		return
	}

	// 更新数据库
	userInfo.IsStop = "use"
	err = model.UserNew(userInfo).UpdatedByUid()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "数据库保存数据失败，请再次重试！",
			"data": "",
		})
		return
	}

	// 获取系统初始信息
	core.CronDoNew().GetLogin()
	core.CronDoNew().GetEveryDayClassTable()
	core.CronDoNew().UpdatePkData()

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}

// 验证参数
func installVerifyValidate(studentName, studentID, timestamp, sign string) error {
	if len(strings.TrimSpace(studentName)) == 0 {
		return errors.New("真实姓名不能为空")
	}

	if len(strings.TrimSpace(studentName)) > 20 {
		return errors.New("真实姓名太长，正常人没有这么长的姓名吧！")
	}

	if len(strings.TrimSpace(studentID)) == 0 {
		return errors.New("学号不能为空")
	}

	if len(strings.TrimSpace(studentID)) != 8 {
		return errors.New("学号不存在，请确认后再次输入！")
	}

	tempStudentID, _ := strconv.Atoi(studentID[:4])
	if tempStudentID < time.Now().Year()-3 {
		return errors.New("学号输入不正确，您已毕业无法使用该系统！")
	} else if tempStudentID == time.Now().Year() && time.Now().Month() < 9 {
		return errors.New("您还未入学,请入学后再使用该系统！")
	} else if tempStudentID > time.Now().Year() {
		return errors.New("学号不存在，请确认后再次输入！")
	}

	err := utils.SafeApp.SafeVerify(gin.H{
		"StudentName": studentName,
		"StudentID":   studentID,
		"timestamp":   timestamp,
		"sign":        sign,
	})
	if err != nil {
		return err
	}

	return nil
}
