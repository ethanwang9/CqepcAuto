package service

import (
	"CqepcAutoServer/service/internal"
	"github.com/robfig/cron/v3"
)

// name: 定时任务
// author: Ethan.Wang
// desc:

// InitCron 初始化 Cron
func InitCron() *cron.Cron {
	return cron.New(cron.WithSeconds())
}

// CronRun 启动定时任务
func CronRun() {
	internal.CronToDo.TODO().Run()
}
