package global

// name: 全局常量
// author: Axel Wong

// 应用信息

const (
	// Version 版本号
	Version = "1.0.6"
	// Author 作者
	Author = "Ethan.Wang"
	// AppUrl 项目地址
	AppUrl = "https://github.com/ethanwang9/CqepcAuto"
)

// 应用配置

const (
	// Mode 模式
	Mode = "release"
	// Port 端口号
	Port = "34567"
)

// 日志

const (
	Director      = "log"
	ShowLine      = true
	StacktraceKey = "stacktrace"
	EncodeLevel   = "CapitalLevelEncoder"
	Format        = "console"
	Prefix        = "[CqepcAuto]"
	LogInConsole  = false
)

// 接口

const (
	Cqepc     = "https://ls.smrte.cn"
	Referer   = "https://servicewechat.com/wx40d0e8c8e2f058d6/25/page-frame.html"
	UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"
)

// 安全

const (
	SafeKey = `7o5cbh0TwOt4gB2y`
	WebKey  = `f9689642a013449eb2cf28bb928dbc87`
)
