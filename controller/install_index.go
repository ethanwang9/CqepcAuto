package controller

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/axelwong/CqepcAuto/api"
	"github.com/axelwong/CqepcAuto/api/cqepc"
	"github.com/axelwong/CqepcAuto/global"
	"github.com/axelwong/CqepcAuto/model"
	"github.com/axelwong/CqepcAuto/utils"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func InstallIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "install/index.html", gin.H{
		"version": global.Version,
	})
}

func InstallIndexPost(ctx *gin.Context) {
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

	// 写入数据库
	index := strings.Index(strings.TrimSpace(dingTalkWebhook), "access_token") + len("access_token") + 1
	token := dingTalkWebhook[index:]
	err = model.UserNew(model.User{
		Uid:        utils.ToolApp.UUID(),
		SId:        userInfo.Data.Username,
		SPass:      strings.TrimSpace(studentPassword),
		SOpenid:    strings.TrimSpace(studentOpenid),
		SName:      userInfo.Data.Nickname,
		SClass:     userInfo.Data.ClassName,
		SClassCode: userInfo.Data.ClassCode,
		PkToken:    userInfo.Data.Token,
		LoginType:  strings.TrimSpace(sysLoginType),
		NToken:     token,
		NSecret:    strings.TrimSpace(dingTalkSecret),
		NPhone:     strings.TrimSpace(dingTalkPhone),
		IsStop:     "register",
	}).Add()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "写入信息到数据库失败，请再次重试！",
			"data": "",
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}

// 校验规则
func installIndexValidate(ctx *gin.Context, studentID, studentPassword, studentOpenid, dingTalkWebhook, dingTalkSecret, sysLoginType, sysCode, sign, timestamp, dingTalkPhone string) error {
	// 校验学号
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

	// 校验密码
	if len(strings.TrimSpace(studentPassword)) == 0 {
		return errors.New("密码不能为空")
	}

	// 消息通知校验
	if len(strings.TrimSpace(dingTalkWebhook)) == 0 {
		return errors.New("消息通知Webhook地址不能为空")
	}

	if len(strings.TrimSpace(dingTalkSecret)) == 0 {
		return errors.New("消息通知密钥不能为空")
	}

	if len(strings.TrimSpace(dingTalkPhone)) == 0 {
		return errors.New("手机号不能为空")
	}

	if b, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", dingTalkPhone); !b {
		return errors.New("请输入正确的手机号")
	}

	// 登录类型校验
	if len(strings.TrimSpace(sysLoginType)) == 0 {
		return errors.New("登录类型不能为空")
	}

	if strings.TrimSpace(sysLoginType) != "account" && strings.TrimSpace(sysLoginType) != "openid" {
		return errors.New("登录类型错误，请检查配置后再提交")
	}

	if strings.TrimSpace(sysLoginType) == "openid" && len(strings.TrimSpace(studentOpenid)) == 0 {
		return errors.New("您当前的登录模式为：评课小程序OPENID登录，评课小程序OPENID不能为空")
	}

	// 校验验证码
	if len(strings.TrimSpace(sysCode)) == 0 {
		return errors.New("验证码不能为空")
	}

	// 校验sign
	if len(strings.TrimSpace(sign)) == 0 {
		return errors.New("浏览器不安全，请检查后重试")
	}

	err := utils.SafeApp.SafeVerify(map[string]interface{}{
		"StudentID":       studentID,
		"StudentPassword": studentPassword,
		"StudentOpenid":   studentOpenid,
		"DingTalkWebhook": dingTalkWebhook,
		"DingTalkSecret":  dingTalkSecret,
		"DingTalkPhone":   dingTalkPhone,
		"SysLoginType":    sysLoginType,
		"SysCode":         sysCode,
		"sign":            sign,
		"timestamp":       timestamp,
	})
	if err != nil {
		return err
	}

	// 校验验证码
	verifyCode, _ := ctx.Cookie("code")
	iv, _ := ctx.Cookie("iv")

	if len(verifyCode) == 0 || len(iv) == 0 {
		return errors.New("请获取验证码")
	}

	code, _ := hex.DecodeString(verifyCode)
	iv2, _ := strconv.ParseInt(iv, 10, 64)
	codeCore, err := goEncrypt.AesCbcDecrypt(code, []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv2-1)))
	if err != nil {
		return errors.New("浏览器不安全，验证码解析失败！ Error: " + err.Error())
	}

	if string(codeCore) != strings.TrimSpace(sysCode) {
		return errors.New("验证码不正确")
	}

	return nil
}
