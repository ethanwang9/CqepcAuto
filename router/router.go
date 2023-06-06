package router

import (
	"github.com/axelwong/CqepcAuto/global"
	"github.com/axelwong/CqepcAuto/middleware"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	// gin 模式
	if global.Mode == "debug" {
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
	router.Static("/static", "./static")

	// 模板渲染
	router.LoadHTMLGlob("./view/**/*")

	// 路由
	r := router.Group("/")
	{
		// 默认
		V1ByInit(r)
		// 安装
		V1ByInstall(r)
		// 管理
		V1ByAdmin(r)
		// 接口
		V1ByApi(r)
	}

	return router
}
