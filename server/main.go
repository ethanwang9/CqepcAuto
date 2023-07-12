package main

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/router"
	"CqepcAutoServer/service"
	"fmt"
	"go.uber.org/zap"
)

// name: 程序入口
// author: Ethan.Wang
// desc: 初始化项目服务

func main() {
	// 初始化配置文件
	global.CONFIG = service.InitViper()
	// 初始化日志
	global.LOG = service.InitZap()
	// 初始化数据库
	// TODO 数据库表结构需要优化
	global.DB = service.InitDB()
	// 初始化定时任务
	global.CRON = service.InitCron()
	defer global.CRON.Stop()
	// 初始化路由
	r := router.Init()

	// 启动定时任务
	service.CronRun()

	// 运行
	err := r.Run(fmt.Sprintf(":%v", global.CONFIG.GetString("server.port")))
	if err != nil {
		global.LOG.Error("系统运行失败", zap.String("info", err.Error()))
		return
	} else {
		global.LOG.Info("系统已开始运行", zap.String("port", global.CONFIG.GetString("server.port")))
	}
}
