package cqepc

import (
	"encoding/json"
	"fmt"
	"github.com/axelwong/CqepcAuto/global"
	"github.com/axelwong/CqepcAuto/utils"
	"github.com/go-resty/resty/v2"
	"github.com/wumansgy/goEncrypt"
	"go.uber.org/zap"
	"time"
)

// ClassPK 评课数据
type ClassPK struct {
	Token       string `json:"token"`
	ClassroomId string `json:"classroom_id"`
}

// New 初始化
func (c *ClassPK) New(classPK ClassPK) *ClassPK {
	return &classPK
}

// Get 获取
func (c *ClassPK) Get() (global.ResPKGet, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/extends/getLessonsInfo"

	var data global.ResPKGet

	// 发起请求
	client := resty.New()
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": global.UserAgent,
			"Referer":    global.Referer,
			"token":      c.Token,
		}).
		SetBody(map[string]string{
			"classroomId": c.ClassroomId,
		}).
		SetResult(&data).
		Post(url)

	if err != nil {
		global.APP_LOG.Error("请求 cqepc#评课数据 失败", zap.Error(err))
		return global.ResPKGet{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := goEncrypt.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#评课数据 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}
