package api

import (
	"CqepcAuto/api/cqepc"
	"CqepcAuto/api/dingtalk"
)

// name: 接口集合
// author: axel wong
// desc:

// ApiGroup 接口集合
type ApiGroup struct {
	CqepcGroup    cqepc.CqepcGroup
	DingTalkGroup dingtalk.DingTalkGroup
}

var ApiGroupApp = new(ApiGroup)
