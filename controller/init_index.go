package controller

import (
	"CqepcAuto/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitIndex(ctx *gin.Context) {
	// 判断系统是否已初始化
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.String(http.StatusOK, "数据库连接异常，Error: "+err.Error())
		return
	}

	if len(userInfo.Uid) == 0 && len(userInfo.IsStop) == 0 {
		// 未初始化
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/install/index?t=%v", time.Now().Unix()))
		return
	} else if len(userInfo.Uid) != 0 && userInfo.IsStop == "register" {
		// 未激活用户
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/install/verify?t=%v", time.Now().Unix()))
		return
	} else if len(userInfo.Uid) != 0 && userInfo.IsStop == "use" {
		// 系统已初始化
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/index?t=%v", time.Now().Unix()))
		return
	} else {
		ctx.JSON(http.StatusOK, "系统配置出现错误，建议重新安装系统！")
	}
}
