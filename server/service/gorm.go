package service

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/model"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// name: GORM 数据库
// author: Ethan.Wang
// desc: sqlite3 数据库初始化

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表名复数
		},
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(global.DBFile), &gormConfig)
	if err != nil {
		global.LOG.Panic("数据库初始化失败", zap.Error(err))
	}

	// 连接数据库池
	dbs, err := db.DB()
	if err != nil {
		global.LOG.Panic("数据库连接池初始化错误", zap.Error(err))
	}
	// 设置空闲连接池中的最大连接数
	dbs.SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数
	dbs.SetMaxOpenConns(100)
	// 设置连接可重用的最大时间
	dbs.SetConnMaxLifetime(time.Hour)

	// 自动注册数据库模型
	autoCreateTable(db)

	return db
}

// 批量注册数据库模型
func autoCreateTable(db *gorm.DB) {
	tables := []interface{}{
		model.ClassInfo{},
		model.ClassToday{},
		model.SysBase{},
		model.SysLog{},
		model.UserAuth{},
		model.UserBase{},
		model.UserDetails{},
		model.UserMessage{},
	}
	err := db.AutoMigrate(tables...)
	if err != nil {
		global.LOG.Panic("数据库初始化-批量注册数据库模型失败", zap.Error(err))
		return
	}
}
