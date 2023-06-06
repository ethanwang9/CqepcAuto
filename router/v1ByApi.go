package router

import (
	"github.com/axelwong/CqepcAuto/controller"
	"github.com/gin-gonic/gin"
)

func V1ByApi(group *gin.RouterGroup) {
	api := group.Group("/api")
	{
		// 获取验证码 - 安装系统初始化
		api.POST("/code/install", controller.ApiCode)
		// 获取验证码 - 更新系统配置
		api.POST("/code/update", controller.ApiSysUpdate)
		// 补评课程
		api.POST("/pk/assist", controller.ApiPkAssist)
	}
}
