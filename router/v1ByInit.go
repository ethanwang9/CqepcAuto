package router

import (
	"github.com/axelwong/CqepcAuto/controller"
	"github.com/gin-gonic/gin"
)

func V1ByInit(group *gin.RouterGroup) {
	api := group.Group("/")
	{
		// 默认页面
		api.GET("/", controller.InitIndex)
	}
}
