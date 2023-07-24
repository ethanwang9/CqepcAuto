package router

import (
	"CqepcAuto/controller"
	"CqepcAuto/middleware"
	"github.com/gin-gonic/gin"
)

func V1ByAdmin(group *gin.RouterGroup) {
	admin := group.Group("/admin")
	admin.Use(middleware.VerifyInstalled())
	{
		admin.GET("/", controller.AdminIndex)
		// 首页
		admin.GET("/index", controller.AdminIndex)
		// 系统配置
		admin.GET("/system", controller.AdminSystem)
		admin.POST("/system", controller.AdminSystemPost)
	}
}
