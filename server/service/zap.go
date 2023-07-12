package service

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/service/internal"
	"CqepcAutoServer/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// name: 写入日志
// author: GVA [github.com/flipped-aurora/gin-vue-admin]
// desc: 写入系统日志库

// InitZap 初始化日志
func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.CONFIG.GetString("log.dir")); !ok { // 判断是否有Director文件夹
		fmt.Printf("初始化日志，创建日志目录: %v\n", global.CONFIG.GetString("log.dir"))
		_ = os.Mkdir(global.CONFIG.GetString("log.dir"), os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.GetBool("log.show_line") {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
