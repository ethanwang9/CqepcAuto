package global

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// name: 全局变量
// author: Axel Wong

var (
	// APP_LOG 日志
	APP_LOG *zap.Logger
	// APP_CRON cron 定时器
	APP_CRON *cron.Cron
	// APP_DB 数据库
	APP_DB *gorm.DB
)
