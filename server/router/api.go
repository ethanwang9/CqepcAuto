package router

import "github.com/gin-gonic/gin"

// name: 路由-三方接口
// author: Ethan.Wang
// desc: 系统使用的其他接口

func Api(r *gin.RouterGroup) *gin.RouterGroup {
	a := r.Group("/api")

	return a
}
