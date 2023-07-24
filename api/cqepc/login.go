package cqepc

import (
	"CqepcAuto/global"
	"CqepcAuto/utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"time"
)

// 登录

type Login struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Openid   string `json:"openid"`
}

// New 初始化
func (l *Login) New(login Login) *Login {
	return &login
}

// UPLogin 账号密码登录
func (l *Login) UPLogin() (global.ResLogin, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/login"

	if len(l.Openid) == 0 {
		l.Openid = utils.ToolApp.RandWxOpenid()
	}

	var data global.ResLogin

	// 发起请求
	client := resty.New()
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": global.UserAgent,
			"Referer":    global.Referer,
			"token":      utils.ToolApp.RandString(32, 2),
		}).
		SetBody(l).
		SetResult(&data).
		Post(url)

	if err != nil {
		global.APP_LOG.Error("请求 cqepc#用户名密码登录 失败", zap.Error(err))
		return global.ResLogin{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := aes.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#用户名密码登录 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}

// AutoLogin 自动登录
func (l *Login) AutoLogin() (global.ResLogin, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/loginOpenid"

	var data global.ResLogin

	// 发起请求
	client := resty.New()
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": global.UserAgent,
			"Referer":    global.Referer,
			"token":      utils.ToolApp.RandString(32, 2),
		}).
		SetBody("{}").
		SetQueryParam("openid", l.Openid).
		SetResult(&data).
		Post(url)

	if err != nil {
		global.APP_LOG.Error("请求 cqepc#自动登录 失败", zap.Error(err))
		return global.ResLogin{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := aes.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#自动登录 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}
