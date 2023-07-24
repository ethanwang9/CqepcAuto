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

// ClassToday 今日课程
type ClassToday struct {
	Token string `json:"token"`
}

// New 初始化
func (c *ClassToday) New(classToday ClassToday) *ClassToday {
	return &classToday
}

// Get 获取
func (c *ClassToday) Get() (global.ResTodayClass, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/queryClassroomForStudent"

	var data global.ResTodayClass

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
		global.APP_LOG.Error("请求 cqepc#今日课程 失败", zap.Error(err))
		return global.ResTodayClass{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := aes.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#今日课程 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}
