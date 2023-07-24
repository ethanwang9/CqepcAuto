package middleware

import (
	"CqepcAuto/global"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
	"time"
)

// name: 处理日志
// author: axel wong
// desc: 中间件处理日志，记录http请求

// LogLayout 日志layout
type LogLayout struct {
	Method string              // 请求类型
	Path   string              // 访问路径
	Query  string              // Get请求参数
	Header map[string][]string // 请求头
	Body   string              // 请求body参数
	IP     string              // ip地址
	Error  string              // 错误
	Cost   time.Duration       // 花费时间
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 访问路径
		path := c.Request.URL.Path
		// Get请求参数
		query := c.Request.URL.RawQuery
		// 请求body参数
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// 请求类型
		method := c.Request.Method

		c.Next()

		// 花费时间
		cost := time.Since(start)
		// 状态码
		status := strconv.Itoa(c.Writer.Status())

		// 写入日志
		global.APP_LOG.Info(
			fmt.Sprintf("%s %s", status, method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Any("header", c.Request.Header),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.ByteString("body", body),
			zap.String("ip", c.ClientIP()),
			zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)

	}
}
