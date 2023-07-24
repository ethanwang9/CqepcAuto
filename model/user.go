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

// User 用户表
type User struct {
	Uid           string `gorm:"column:uid" json:"uid,omitempty"`
	SId           string `gorm:"column:s_id" json:"s_id,omitempty"`
	SPass         string `gorm:"column:s_pass" json:"s_pass,omitempty"`
	SOpenid       string `gorm:"column:s_openid" json:"s_openid,omitempty"`
	SName         string `gorm:"column:s_name" json:"s_name,omitempty"`
	SClass        string `gorm:"column:s_class" json:"s_class,omitempty"`
	SClassCode    string `gorm:"column:s_class_code" json:"s_class_code,omitempty"`
	PkToken       string `gorm:"column:pk_token" json:"pk_token,omitempty"`
	LoginType     string `gorm:"column:login_type" json:"login_type,omitempty"`
	NToken        string `gorm:"column:n_token" json:"n_token,omitempty"`
	NSecret       string `gorm:"column:n_secret" json:"n_secret,omitempty"`
	NPhone        string `gorm:"column:n_phone" json:"n_phone,omitempty"`
	IsStop        string `gorm:"column:is_stop" json:"is_stop,omitempty"`
	CqepcAutoFlag string `gorm:"column:cqepc_auto_flag" json:"cqepc_auto_flag,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// UserNew 初始化
func UserNew(user User) *User {
	return &user
}

// Add 添加用户
func (u *User) Add() error {
	userEncode(u)

	err := global.APP_DB.Create(&u).Error
	if err != nil {
		global.APP_LOG.Error("数据库-user表-添加用户失败", zap.Error(err))
		return err
	}

	return nil
}

// GetByUid 通过uid查询用户信息
func (u *User) GetByUid() (User, error) {
	var user User
	err := global.APP_DB.Where("uid = ?", u.Uid).Find(&user).Error
	if err != nil {
		global.APP_LOG.Error("数据库-user表-通过uid查询用户信息失败", zap.Error(err))
		return user, err
	}

	userDecode(&user)

	return user, nil
}

// UpdatedByUid 通过用户uid更新用户信息
func (u *User) UpdatedByUid() error {
	userEncode(u)

	err := global.APP_DB.Where("uid = ?", u.Uid).Updates(&u).Error
	if err != nil {
		global.APP_LOG.Error("数据库-user表-通过用户uid更新用户信息失败", zap.Error(err))
		return err
	}
	return nil
}

// GetByUidOldest 获取最老的用户ID
func (u *User) GetByUidOldest() (User, error) {
	var user User
	err := global.APP_DB.Order("created_at").Limit(1).Find(&user).Error
	if err != nil {
		global.APP_LOG.Error("数据库-user表-获取最老的用户ID失败", zap.Error(err))
		return user, err
	}

	userDecode(&user)

	return user, nil
}

// 加密
func userEncode(user *User) {
	iv := time.Now().UnixMicro()

	sid, _ := aes.AesCbcEncrypt([]byte(user.SId), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SId = hex.EncodeToString(sid)

	sPass, _ := aes.AesCbcEncrypt([]byte(user.SPass), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SPass = hex.EncodeToString(sPass)

	sOpenid, _ := aes.AesCbcEncrypt([]byte(user.SOpenid), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SOpenid = hex.EncodeToString(sOpenid)

	sName, _ := aes.AesCbcEncrypt([]byte(user.SName), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SName = hex.EncodeToString(sName)

	sClass, _ := aes.AesCbcEncrypt([]byte(user.SClass), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SClass = hex.EncodeToString(sClass)

	sClassCode, _ := aes.AesCbcEncrypt([]byte(user.SClassCode), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SClassCode = hex.EncodeToString(sClassCode)

	pkToken, _ := aes.AesCbcEncrypt([]byte(user.PkToken), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.PkToken = hex.EncodeToString(pkToken)

	loginType, _ := aes.AesCbcEncrypt([]byte(user.LoginType), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.LoginType = hex.EncodeToString(loginType)

	nToken, _ := aes.AesCbcEncrypt([]byte(user.NToken), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NToken = hex.EncodeToString(nToken)

	nSecret, _ := aes.AesCbcEncrypt([]byte(user.NSecret), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NSecret = hex.EncodeToString(nSecret)

	nPhone, _ := aes.AesCbcEncrypt([]byte(user.NPhone), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NPhone = hex.EncodeToString(nPhone)

	isStop, _ := aes.AesCbcEncrypt([]byte(user.IsStop), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.IsStop = hex.EncodeToString(isStop)

	user.CqepcAutoFlag = fmt.Sprintf("%v", iv+1)
}

// 解密
func userDecode(user *User) {
	iv, _ := strconv.ParseInt(user.CqepcAutoFlag, 10, 64)
	iv--

	sid, _ := hex.DecodeString(user.SId)
	sid, _ = aes.AesCbcDecrypt([]byte(sid), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SId = string(sid)

	sPass, _ := hex.DecodeString(user.SPass)
	sPass, _ = aes.AesCbcDecrypt([]byte(sPass), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SPass = string(sPass)

	sOpenid, _ := hex.DecodeString(user.SOpenid)
	sOpenid, _ = aes.AesCbcDecrypt([]byte(sOpenid), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SOpenid = string(sOpenid)

	sName, _ := hex.DecodeString(user.SName)
	sName, _ = aes.AesCbcDecrypt([]byte(sName), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SName = string(sName)

	sClass, _ := hex.DecodeString(user.SClass)
	sClass, _ = aes.AesCbcDecrypt([]byte(sClass), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SClass = string(sClass)

	sClassCode, _ := hex.DecodeString(user.SClassCode)
	sClassCode, _ = aes.AesCbcDecrypt([]byte(sClassCode), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.SClassCode = string(sClassCode)

	pkToken, _ := hex.DecodeString(user.PkToken)
	pkToken, _ = aes.AesCbcDecrypt([]byte(pkToken), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.PkToken = string(pkToken)

	loginType, _ := hex.DecodeString(user.LoginType)
	loginType, _ = aes.AesCbcDecrypt([]byte(loginType), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.LoginType = string(loginType)

	nToken, _ := hex.DecodeString(user.NToken)
	nToken, _ = aes.AesCbcDecrypt([]byte(nToken), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NToken = string(nToken)

	nSecret, _ := hex.DecodeString(user.NSecret)
	nSecret, _ = aes.AesCbcDecrypt([]byte(nSecret), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NSecret = string(nSecret)

	nPhone, _ := hex.DecodeString(user.NPhone)
	nPhone, _ = aes.AesCbcDecrypt([]byte(nPhone), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.NPhone = string(nPhone)

	isStop, _ := hex.DecodeString(user.IsStop)
	isStop, _ = aes.AesCbcDecrypt([]byte(isStop), []byte(global.SafeKey), []byte(fmt.Sprintf("%v", iv)))
	user.IsStop = string(isStop)

	user.CqepcAutoFlag = fmt.Sprintf("%v", iv)
}
