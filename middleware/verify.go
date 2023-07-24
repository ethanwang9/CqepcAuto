package middleware

import (
	"CqepcAuto/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// VerifyInstalled 判断是否已安装系统
func VerifyInstalled() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
		if err != nil {
			ctx.String(http.StatusOK, "读取数据库失败，请再次重试！ Error: "+err.Error())
			ctx.Abort()
			return
		}

		if len(userInfo.Uid) == 0 && len(userInfo.IsStop) == 0 {
			if path != "/install/index" {
				ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/install/index?t=%v", time.Now().Unix()))
				ctx.Abort()
				return
			}
		} else if len(userInfo.Uid) != 0 && userInfo.IsStop == "register" {
			// 未激活用户
			if path != "/install/verify" {
				ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/install/verify?t=%v", time.Now().Unix()))
				ctx.Abort()
				return
			}
		} else if len(userInfo.Uid) != 0 && userInfo.IsStop == "use" {
			// 系统已初始化
			if path == "/admin/" {
				ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/index?t=%v", time.Now().Unix()))
				ctx.Abort()
				return
			} else if !strings.HasPrefix(path, "/admin/") {
				ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/index?t=%v", time.Now().Unix()))
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusOK, "系统配置出现错误，建议重新安装系统！")
			ctx.Abort()
			return
		}

		// end
	}
}
