package model

import (
	"CqepcAuto/global"
	"CqepcAuto/utils"
	"encoding/hex"
	"fmt"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// Today 今日课表
type Today struct {
	Uid           string `gorm:"column:uid" json:"uid,omitempty"`
	PId           string `gorm:"column:p_id" json:"p_id,omitempty"`
	Data          string `gorm:"column:data" json:"data,omitempty"`
	PkData        string `gorm:"column:pk_data" json:"pk_data,omitempty"`
	IsPk          string `gorm:"column:is_pk" json:"is_pk,omitempty"`
	CqepcAutoFlag string `gorm:"column:cqepc_auto_flag" json:"cqepc_auto_flag,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TodayNew 初始化
func TodayNew(t Today) *Today {
	return &t
}

// Add 添加
func (t *Today) Add() error {
	todayEncode(t)

	err := global.APP_DB.Create(&t).Error
	if err != nil {
		global.APP_LOG.Error("数据库-今日课表-添加今日课表失败", zap.Error(err))
		return err
	}

	return nil
}

// GetByUidToday 通过uid获取今日最新的课表
func (t *Today) GetByUidToday() ([]Today, error) {
	var today []Today
	t.CreatedAt = utils.ToolApp.TodayZeroUnix()
	err := global.APP_DB.Where("uid = ? and created_at >= ?", t.Uid, t.CreatedAt).Find(&today).Error
	if err != nil {
		global.APP_LOG.Error("数据库-today表-通过uid获取今日最新的课表失败", zap.Error(err))
		return []Today{}, err
	}

	core := todayDecode(today)

	return core, nil
}

// GetByUidAndPID 获取用户ID和课程id获取课程信息
func (t *Today) GetByUidAndPID() (Today, error) {
	var today []Today
	err := global.APP_DB.Where("uid = ? and p_id = ?", t.Uid, t.PId).Find(&today).Error
	if err != nil {
		global.APP_LOG.Error("数据库-today表-获取用户ID和课程id获取课程信息", zap.Error(err))
		return Today{}, err
	}

	core := todayDecode(today)

	return core[0], nil
}

// DelTodayClassByUid 通过UID删除今天课表
func (t *Today) DelTodayClassByUid() error {
	t.CreatedAt = utils.ToolApp.TodayZeroUnix()
	err := global.APP_DB.Where("uid = ? and created_at >= ?", t.Uid, t.CreatedAt).Delete(&t).Error
	if err != nil {
		global.APP_LOG.Error("数据库-today表-通过UID删除今天课表失败", zap.Error(err))
		return err
	}

	return nil
}

// UpdateByUidAndPid 通过uid和pid更新课表状态
func (t *Today) UpdateByUidAndPid() error {
	todayEncode(t)
	err := global.APP_DB.Where("uid = ? and p_id = ?", t.Uid, t.PId).Updates(&t).Error
	if err != nil {
		global.APP_LOG.Error("数据库-today表-通过uid和pid更新课表状态失败", zap.Error(err))
		return err
	}

	return nil
}

// 加密
func todayEncode(t *Today) {
	iv := time.Now().UnixMicro()

	data, _ := aes.AesCbcEncrypt([]byte(t.Data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	t.Data = hex.EncodeToString(data)

	pkData, _ := aes.AesCbcEncrypt([]byte(t.PkData), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	t.PkData = hex.EncodeToString(pkData)

	isPk, _ := aes.AesCbcEncrypt([]byte(t.IsPk), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	t.IsPk = hex.EncodeToString(isPk)

	t.CqepcAutoFlag = fmt.Sprintf("%v", iv+1)
}

// 解密
func todayDecode(t []Today) []Today {
	core := make([]Today, 0)
	for _, v := range t {
		iv, _ := strconv.ParseInt(v.CqepcAutoFlag, 10, 64)
		iv--

		data, _ := hex.DecodeString(v.Data)
		data, _ = aes.AesCbcDecrypt([]byte(data), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.Data = string(data)

		pkData, _ := hex.DecodeString(v.PkData)
		pkData, _ = aes.AesCbcDecrypt([]byte(pkData), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.PkData = string(pkData)

		isPk, _ := hex.DecodeString(v.IsPk)
		isPk, _ = aes.AesCbcDecrypt([]byte(isPk), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
		v.IsPk = string(isPk)

		v.CqepcAutoFlag = fmt.Sprintf("%v", iv)

		core = append(core, v)
	}

	return core
}
