package global

// name: 请求返回数据结构
// author: Ethan.Wang
// desc:

// MsgBack 消息返回结构体
type MsgBack struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
