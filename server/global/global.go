package global

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// name: 全局变量
// author: Ethan.Wang
// desc: 程序运行主要变量

var (
	CONFIG *viper.Viper
	LOG    *zap.Logger
	CRON   *cron.Cron
	DB     *gorm.DB
)
