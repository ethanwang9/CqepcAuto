package main

import (
	"CqepcAuto/core"
	"CqepcAuto/global"
	"CqepcAuto/router"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 初始化日志
	global.APP_LOG = core.Zap()
	// 初始化数据库
	global.APP_DB = core.GormBySqliteInit()
	// 初始化定时任务
	global.APP_CRON = core.CronInit()
	defer global.APP_CRON.Stop()
	// 初始化路由
	r := router.Init()

	// 启动定时任务
	core.CronDoNew().DO().Run()

	// 写入日志
	global.APP_LOG.Info("CqepcAuto 服务成功运行", zap.String("port", global.Port))

	// 显示控制台日志
	consoleLog()

	// 运行
	_ = r.Run(fmt.Sprintf(":%v", global.Port))
}

func consoleLog() {
	logo := " ____                                    ______           __             \n/\\  _`\\                                 /\\  _  \\         /\\ \\__          \n\\ \\ \\/\\_\\     __      __   _____     ___\\ \\ \\L\\ \\  __  __\\ \\ ,_\\   ___   \n \\ \\ \\/_/_  /'__`\\  /'__`\\/\\ '__`\\  /'___\\ \\  __ \\/\\ \\/\\ \\\\ \\ \\/  / __`\\ \n  \\ \\ \\L\\ \\/\\ \\L\\ \\/\\  __/\\ \\ \\L\\ \\/\\ \\__/\\ \\ \\/\\ \\ \\ \\_\\ \\\\ \\ \\_/\\ \\L\\ \\\n   \\ \\____/\\ \\___, \\ \\____\\\\ \\ ,__/\\ \\____\\\\ \\_\\ \\_\\ \\____/ \\ \\__\\ \\____/\n    \\/___/  \\/___/\\ \\/____/ \\ \\ \\/  \\/____/ \\/_/\\/_/\\/___/   \\/__/\\/___/ \n                 \\ \\_\\       \\ \\_\\                                       \n                  \\/_/        \\/_/                                       "
	fmt.Println(logo)
	fmt.Println("\n>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(fmt.Sprintf("请访问面板地址\thttp://127.0.0.1:%v/?t=%v", global.Port, time.Now().Unix()))
}
