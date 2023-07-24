package model

import (
	"CqepcAuto/global"
	"CqepcAuto/utils"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// Log 日志表
type Log struct {
	Uuid          string `gorm:"column:uuid" json:"uuid,omitempty"`
	Uid           string `gorm:"column:uid" json:"uid,omitempty"`
	Msg           string `gorm:"column:msg" json:"msg,omitempty"`
	MsgType       string `gorm:"column:msg_type" json:"msg_type,omitempty"`
	Data          string `gorm:"column:data" json:"data,omitempty"`
	CqepcAutoFlag string `gorm:"column:cqepc_auto_flag" json:"cqepc_auto_flag,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// LogNew 初始化
func LogNew(log Log) *Log {
	return &log
}

// Add 添加
func (l *Log) Add() error {
	logEncode(l)

	err := global.APP_DB.Create(&l).Error
	if err != nil {
		global.APP_LOG.Error("数据库-log表-添加日志失败", zap.Error(err))
		return err
	}

	return nil
}

// GetByUid 通过用户UID获取数据信息
func (l *Log) GetByUid() ([]Log, error) {
	var log []Log
	err := global.APP_DB.Where("uid = ? and msg_type = ? and created_at >= ?", l.Uid, "ding_talk_msg", l.CreatedAt).Find(&log).Error
	if err != nil {
		global.APP_LOG.Error("数据库-log表-通过用户UID获取数据信息失败", zap.Error(err))
		return []Log{}, err
	}

	core := logDecode(log)

	return core, nil
}

// IsSendDingTalk 是否可以发送钉钉消息
// 是 - true
// 否 - false
func (l *Log) IsSendDingTalk() (bool, error) {
	var log []Log
	l.CreatedAt = time.Unix(utils.ToolApp.DateToUnix(fmt.Sprintf("%v-%v-%v %v:%v", time.Now().Year(), time.Now().Format("01"), time.Now().Day(), time.Now().Hour(), time.Now().Minute())), 0)
	err := global.APP_DB.Where("uid = ? and created_at = ?", l.Uid, l.CreatedAt).Find(&log).Error
	if err != nil {
		global.APP_LOG.Error("数据库-log表-是否可以发送钉钉消息获取数据失败", zap.Error(err))
		return false, err
	}

	if len(log) >= 19 {
		return false, errors.New("钉钉发信太频繁，请1分钟后再试！")
	}

	return true, err
}

// 加密
func logEncode(log *Log) {
	iv := time.Now().UnixMicro()

	msg, _ := aes.AesCbcEncrypt([]byte(log.Msg), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	log.Msg = hex.EncodeToString(msg)

	data, _ := aes.AesCbcEncrypt([]byte(log.Data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	log.Data = hex.EncodeToString(data)

	log.CqepcAutoFlag = fmt.Sprintf("%v", iv+1)
}

// 解密
func logDecode(log []Log) []Log {
	core := make([]Log, 0)
	for _, v := range log {
		iv, _ := strconv.ParseInt(v.CqepcAutoFlag, 10, 64)
		iv--

		msg, _ := hex.DecodeString(v.Msg)
		msg, _ = aes.AesCbcDecrypt([]byte(msg), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.Msg = string(msg)

		data, _ := hex.DecodeString(v.Data)
		data, _ = aes.AesCbcDecrypt([]byte(data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.Data = string(data)

		v.CqepcAutoFlag = fmt.Sprintf("%v", iv)

		core = append(core, v)
	}

	return core
}
