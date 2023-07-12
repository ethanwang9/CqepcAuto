package global

// name: 全局常量
// author: Ethan.Wang
// desc:

// 配置文件

const (
	// ConfigFile 配置文件
	ConfigFile = "config.yaml"
	// DBFile 数据库文件
	DBFile = "./ca.db"
)

// 请求状态码 - 成功
const (
	// CodeSuccess 请求成功
	CodeSuccess = 200
)

// 请求状态码 - 系统
const ()

// 请求状态码 - 数据库

const (
	// CodeDBByGetError 获取数据库信息失败
	CodeDBByGetError = 401
)

// 请求状态码 - 接口

const (
	// CodeApiBySignError 签名错误
	CodeApiBySignError = 501
)

// 请求状态码 - 特殊编码
const ()

// 程序基本信息
const (
	AppVersion    = "2.0.0"
	AppAuthor     = "Ethan.Wang"
	AppUpdateTime = 1689173546
	AppGithub     = "https://github.com/ethanwang9/CqepcAuto"
)
