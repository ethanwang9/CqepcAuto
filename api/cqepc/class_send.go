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

// ClassSend 发送数据
type ClassSend struct {
	Token             string `json:"token"`
	HasUnderstand     string `json:"has_understand"`
	HasTeaching       string `json:"has_teaching"`
	HasSchoolwork     string `json:"has_schoolwork"`
	Memo              string `json:"memo"`
	ClassroomId       string `json:"classroom_id"`
	CoincidenceDegree string `json:"coincidence_degree"`
}

// New 初始化
func (c *ClassSend) New(classSend ClassSend) *ClassSend {
	return &classSend
}

// Send 发送评课消息
func (c *ClassSend) Send() (global.ResPKSend, error) {
	// 接口地址
	url := global.Cqepc + "/weChat/evaluateCourse"

	var data global.ResPKSend

	// 发起请求
	client := resty.New()
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": global.UserAgent,
			"Referer":    global.Referer,
			"token":      c.Token,
		}).
		SetBody(map[string]string{
			"has_understand":    c.HasUnderstand,
			"hasTeaching":       c.HasTeaching,
			"hasSchoolwork":     c.HasSchoolwork,
			"memo":              c.Memo,
			"classroomId":       c.ClassroomId,
			"coincidenceDegree": c.CoincidenceDegree,
		}).
		SetResult(&data).
		Post(url)

	if err != nil {
		global.APP_LOG.Error("请求 cqepc#发送数据 失败", zap.Error(err))
		return global.ResPKSend{}, err
	}

	// 写入日志
	dataJsonString, _ := json.Marshal(data)
	dataTime := time.Now().UnixMicro()
	dataIV := []byte(fmt.Sprintf("%v", dataTime))
	dataAesByte, _ := goEncrypt.AesCbcEncrypt(dataJsonString, []byte(global.SafeKey), dataIV)
	global.APP_LOG.Debug(
		"请求 cqepc#发送数据 返回数据",
		zap.Any("data", utils.AlgorithmApp.Base64Encode(dataAesByte)),
		zap.Int64("iv", dataTime+1),
	)

	return data, nil
}
