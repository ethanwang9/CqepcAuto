package router

import (
	"github.com/gin-gonic/gin"
)

// name: 路由-安装接口
// author: Ethan.Wang
// desc: 系统安装时进行数据同步和配置保存

func Install(r *gin.RouterGroup) *gin.RouterGroup {
	i := r.Group("/install")

	return i
}
