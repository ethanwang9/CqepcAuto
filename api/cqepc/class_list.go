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

// ClassList 课程列表
type ClassList struct {
	Token string `json:"token"`
}

// New 初始化
func (c *ClassList) New(classList ClassList) *ClassList {
	return &classList
}

// Get 获取
func (c *ClassList) Get() (global.ResClassAll, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/selectClassroomDataList"

	var data global.ResClassAll

	// 发起请求
	client := resty.New()
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": global.UserAgent,
			"Referer":    global.Referer,
			"token":      c.Token,
		}).
		SetBody("{}").
		SetResult(&data).
		Post(url)

	if err != nil {
		global.APP_LOG.Error("请求 cqepc#获取课程列表 失败", zap.Error(err))
		return global.ResClassAll{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := aes.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#获取课程列表 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}
