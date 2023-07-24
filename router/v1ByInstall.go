package router

import (
	"CqepcAuto/controller"
	"CqepcAuto/middleware"
	"github.com/gin-gonic/gin"
)

func V1ByInstall(group *gin.RouterGroup) {
	api := group.Group("/install")
	api.Use(middleware.VerifyInstalled())
	{
		api.GET("/", controller.InstallIndex)

		// 安装页面
		api.GET("/index", controller.InstallIndex)
		api.POST("/index", controller.InstallIndexPost)
		// 校验页面
		api.GET("/verify", controller.InstallVerify)
		api.POST("/verify", controller.InstallVerifyPost)
	}
}
