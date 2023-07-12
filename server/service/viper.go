package service

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/utils"
	"fmt"
	"github.com/spf13/viper"
)

// name: 读取配置
// author: Ethan.Wang
// desc: 读取配置文件数据

// InitViper 初始化 Viper
func InitViper() *viper.Viper {
	// 判断是否有配置文件
	if f, _ := utils.PathExists(global.ConfigFile); !f {
		panic("初始化 [Viper] 失败, 配置文件不存在")
	}

	v := viper.New()
	// 设置配置文件
	v.SetConfigFile(global.ConfigFile)
	// 读取文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("初始化 [Viper] 出现错误, Error: %s\n", err))
	}
	// 监听文件
	v.WatchConfig()
	// 文件改变
	//v.OnConfigChange(func(e fsnotify.Event) {
	//fmt.Println("[Viper] 文件改变", e.Name)
	//})

	return v
}
