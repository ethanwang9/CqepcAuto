package middleware

import (
	"CqepcAutoServer/global"
	"CqepcAutoServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// SafeSign 安全校验 - 验证签名
func SafeSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求类型
		method := c.Request.Method
		// Get 请求参数
		query := c.Request.URL.RawQuery
		// Body 请求参数
		body, _ := c.GetRawData()

		// 解析参数
		var urlMap url.Values
		if method == http.MethodGet {
			urlMap, _ = url.ParseQuery(query)
		} else if method == http.MethodPost {
			urlMap, _ = url.ParseQuery(string(body))
		}

		// 生成 map[string]interface{} 格式数据
		d := make(map[string]interface{}, 0)
		for k, v := range urlMap {
			if len(v) > 0 {
				d[k] = v[0]
			}
		}

		// 验证签名
		if err := utils.SafeApp.SafeVerify(d); err != nil {
			c.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeApiBySignError,
				Message: err.Error(),
				Data:    nil,
			})
			c.Abort()
			return
		}
	}
}
