package router

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/middleware"
	"github.com/gin-gonic/gin"
)

// name: 路由入口
// author: Ethan.Wang
// desc: 加载接口

// Init 初始化路由
func Init() *gin.Engine {
	// gin 模式
	if global.CONFIG.GetString("server.mode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 路由
	router := gin.New()

	// 中间件
	router.Use(middleware.GinLogger())
	router.Use(middleware.GinRecovery(true))

	// 文件目录
	//router.Static("/static", "./static")

	// 模板渲染
	//router.LoadHTMLGlob("./view/**/*")

	// 路由
	v1 := router.Group("/")
	{
		// 安装
		Install(v1)
		// 接口
		Api(v1)
		// 后台
		Admin(v1)
	}

	return router
}
