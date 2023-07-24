package controller

import (
	"CqepcAuto/api"
	"CqepcAuto/api/dingtalk"
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/aes"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// ApiCode 获取验证码 - 安装系统初始化
func ApiCode(ctx *gin.Context) {
	// 获取参数
	token := ctx.PostForm("token")
	secret := ctx.PostForm("secret")
	phone := ctx.PostForm("phone")
	sign := ctx.PostForm("sign")
	timestamp := ctx.PostForm("timestamp")

	// 安全验证
	if err := utils.SafeApp.SafeVerify(gin.H{
		"token":     token,
		"secret":    secret,
		"phone":     phone,
		"sign":      sign,
		"timestamp": timestamp,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求校验未通过，请再次重试！ Error: " + err.Error(),
			"data": "",
		})
		return
	}

	if !strings.HasPrefix(token, "https://oapi.dingtalk.com/robot/send?access_token=") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "消息通知 Webhook 地址不正确，请检查后再试！",
			"data": "",
		})
		return
	}

	if b, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", phone); !b {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请输入正确的手机号",
			"data": "",
		})
		return
	}

	// 处理token
	index := strings.Index(strings.TrimSpace(token), "access_token") + len("access_token") + 1
	token = token[index:]

	// 是否可以发送消息
	isSend, err := model.LogNew(model.Log{Uid: token}).IsSendDingTalk()
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

	// 生成验证码
	code := utils.ToolApp.RandString(6, 0)

	// 生成消息
	msg := []string{
		"# 验证码",
		"---",
		"您正在安装CqepcAuto自动评课系统，现在校验您的钉钉消息配置是否正确！\n",
		"验证码有效期：5分钟",
		"",
		"---",
		"## 验证码：" + code,
	}

	// 写入日志表
	err = model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     token,
		Msg:     "安装系统，验证消息通知是否可以通信",
		MsgType: "ding_talk_msg",
		Data:    "验证码：" + code,
	}).Add()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "处理数据失败，请再次获取验证码",
			"data": "",
		})
		return
	}

	// 发送验证码
	err = api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
		Token:  strings.TrimSpace(token),
		Secret: strings.TrimSpace(secret),
		Phone:  []string{strings.TrimSpace(phone)},
	}).SendMarkdown("验证码", msg)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送验证码失败, Error: " + err.Error(),
			"data": "",
		})
		return
	}

	// 设置cookie
	t := time.Now().UnixMicro()
	codeE, _ := aes.AesCbcEncrypt([]byte(code), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", t)))
	ctx.SetCookie("code", hex.EncodeToString(codeE), 300, "/install", "", false, true)
	ctx.SetCookie("iv", fmt.Sprintf("%v", t+1), 300, "/install", "", false, true)

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取验证码成功，请注意钉钉消息",
		"data": "",
	})
}

// ApiSysUpdate 获取验证码 - 更新系统配置
func ApiSysUpdate(ctx *gin.Context) {
	// 获取参数
	token := ctx.PostForm("token")
	secret := ctx.PostForm("secret")
	phone := ctx.PostForm("phone")
	sign := ctx.PostForm("sign")
	timestamp := ctx.PostForm("timestamp")

	// 安全验证
	if err := utils.SafeApp.SafeVerify(gin.H{
		"token":     token,
		"secret":    secret,
		"phone":     phone,
		"sign":      sign,
		"timestamp": timestamp,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求校验未通过，请再次重试！ Error: " + err.Error(),
			"data": "",
		})
		return
	}

	if !strings.HasPrefix(token, "https://oapi.dingtalk.com/robot/send?access_token=") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "消息通知 Webhook 地址不正确，请检查后再试！",
			"data": "",
		})
		return
	}

	if b, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", phone); !b {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请输入正确的手机号",
			"data": "",
		})
		return
	}

	// 处理token
	index := strings.Index(token, "access_token")
	token = token[index+len("access_token")+1:]

	// 是否可以发送消息
	isSend, err := model.LogNew(model.Log{Uid: token}).IsSendDingTalk()
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

	// 生成验证码
	code := utils.ToolApp.RandString(6, 0)

	// 生成消息
	msg := []string{
		"# 验证码",
		"---",
		"您正在更新CqepcAuto自动评课系统配置，现在校验您的钉钉消息配置是否正确！\n",
		"验证码有效期：5分钟",
		"",
		"---",
		"## 验证码：" + code,
	}

	// 写入日志表
	err = model.LogNew(model.Log{
		Uuid:    utils.ToolApp.UUID(),
		Uid:     strings.TrimSpace(token),
		Msg:     "更新系统配置，验证消息通知是否可以通信",
		MsgType: "ding_talk_msg",
		Data:    "验证码：" + code,
	}).Add()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "处理数据失败，请再次获取验证码",
			"data": "",
		})
		return
	}

	// 发送验证码
	err = api.ApiGroupApp.DingTalkGroup.New(dingtalk.Message{
		Token:  strings.TrimSpace(token),
		Secret: strings.TrimSpace(secret),
		Phone:  []string{strings.TrimSpace(phone)},
	}).SendMarkdown("验证码", msg)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送验证码失败: " + err.Error(),
			"data": "",
		})
		return
	}

	// 设置cookie
	t := time.Now().UnixMicro()
	codeE, _ := aes.AesCbcEncrypt([]byte(code), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", t)))
	ctx.SetCookie("code", hex.EncodeToString(codeE), 300, "/admin/system", "", false, true)
	ctx.SetCookie("iv", fmt.Sprintf("%v", t+1), 300, "/admin/system", "", false, true)

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取验证码成功，请注意钉钉消息",
		"data": "",
	})
}
