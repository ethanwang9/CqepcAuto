package model

import (
	"CqepcAuto/global"
	"encoding/hex"
	"fmt"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// Tj 统计表
type Tj struct {
	Uuid          string `gorm:"column:uuid" json:"uuid,omitempty"`
	Uid           string `gorm:"column:uid" json:"uid,omitempty"`
	Data          string `gorm:"column:data" json:"data,omitempty"`
	CqepcAutoFlag string `gorm:"column:cqepc_auto_flag" json:"cqepc_auto_flag,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TjNew 初始化
func TjNew(tj Tj) *Tj {
	return &tj
}

// Add 添加
func (t *Tj) Add() error {
	tjEncode(t)

	err := global.APP_DB.Create(&t).Error
	if err != nil {
		global.APP_LOG.Error("数据库-统计表-添加统计信息失败", zap.Error(err))
		return err
	}

	return nil
}

// GetByUid 通过uid获取最新统计数据
func (t *Tj) GetByUid() (Tj, error) {
	var tj []Tj
	err := global.APP_DB.Where("uid = ?", t.Uid).Order("created_at DESC").Limit(1).Find(&tj).Error
	if err != nil {
		global.APP_LOG.Error("数据库-tj表-通过uid获取最新统计数据失败", zap.Error(err))
		return Tj{}, err
	}

	core := tjDecode(tj)

	if len(core) == 0 {
		return Tj{}, nil
	} else {
		return core[0], nil
	}
}

// 加密
func tjEncode(tj *Tj) {
	iv := time.Now().UnixMicro()

	data, _ := aes.AesCbcEncrypt([]byte(tj.Data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	tj.Data = hex.EncodeToString(data)

	tj.CqepcAutoFlag = fmt.Sprintf("%v", iv+1)
}

// 解密
func tjDecode(tj []Tj) []Tj {
	core := make([]Tj, 0)
	for _, v := range tj {
		iv, _ := strconv.ParseInt(v.CqepcAutoFlag, 10, 64)
		iv--

		data, _ := hex.DecodeString(v.Data)
		data, _ = aes.AesCbcDecrypt([]byte(data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.Data = string(data)

		v.CqepcAutoFlag = fmt.Sprintf("%v", iv)

		core = append(core, v)
	}

	return core
}
