package cqepc

// name: 评课
// author: axel wong
// desc:

// CqepcGroup 评课系统接口集合
type CqepcGroup struct {
	// 登录
	Login
	// 查看本学期课程表
	ClassList
	// 查看评课详情
	ClassDetails
	// 今日课程
	ClassToday
	// 评课信息
	ClassPK
	// 发送数据
	ClassSend
}
