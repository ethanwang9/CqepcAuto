package router

import "github.com/gin-gonic/gin"

// name: 路由-后台管理系统
// author: Ethan.Wang
// desc: 后台管理系统必要的接口路由

func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	a := r.Group("/admin")

	return a
}
