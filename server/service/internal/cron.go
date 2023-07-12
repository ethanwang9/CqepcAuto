package internal

import (
	"CqepcAutoServer/global"
)

// name: 定时任务
// author: Ethan.Wang
// desc: 定时任务具体执行

type cronToDo struct{}

var CronToDo = new(cronToDo)

// New 初始化实例
func (d *cronToDo) New() *cronToDo {
	return &cronToDo{}
}

// Run 运行
func (d *cronToDo) Run() {
	global.CRON.Start()
}

// TODO 任务列表
func (d *cronToDo) TODO() *cronToDo {
	// 获取微信公众号AccessToken
	//global.CRON.AddFunc("0 0 */1 * * ?", d.GetMpAccessToken)

	return d
}
