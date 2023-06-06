package main

import (
	"encoding/json"
	"github.com/axelwong/CqepcAuto/api"
	"github.com/axelwong/CqepcAuto/api/cqepc"
	"github.com/axelwong/CqepcAuto/core"
	"github.com/axelwong/CqepcAuto/global"
	"github.com/axelwong/CqepcAuto/model"
	"log"
)

func main() {
	// 初始化日志
	global.APP_LOG = core.Zap()
	// 初始化数据库
	global.APP_DB = core.GormBySqliteInit()

	// 修复程序
	fix()
}

func fix() {
	log.Println("==========")
	log.Println("系统修复自检程序")
	log.Println("==========")

	core.CronDoNew().GetLogin()

	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		log.Println("数据库@获取用户信息失败 Error: " + err.Error())
		log.Println("修复程序已被终止")
		return
	}

	// 获取今日课表
	today, err := model.TodayNew(model.Today{Uid: userInfo.Uid}).GetByUidToday()
	if err != nil {
		log.Println("数据库@获取今日课表失败 Error: " + err.Error())
		log.Println("修复程序已被终止")
		return
	}

	if len(today) == 0 {
		cqepcTodayInfo, err := api.ApiGroupApp.CqepcGroup.ClassToday.New(cqepc.ClassToday{Token: userInfo.PkToken}).Get()
		if err != nil {
			log.Println("评课系统@获取今日课表失败 Error: " + err.Error())
			log.Println("修复程序已被终止")
			return
		}
		if cqepcTodayInfo.Code != 200 {
			log.Println("评课系统@获取今日课表失败 Error: " + err.Error())
			log.Println("修复程序已被终止")
			return
		}

		if len(cqepcTodayInfo.Data) != 0 {
			core.CronDoNew().GetEveryDayClassTable()
			log.Println("修复今日课表成功")
		}
	}

	// 获取评课信息
	pk, err := model.TjNew(model.Tj{Uid: userInfo.Uid}).GetByUid()
	if err != nil {
		log.Println("数据库@获取评课信息失败 Error: " + err.Error())
		log.Println("修复程序已被终止")
		return
	}
	if len(pk.Uid) == 0 {
		core.CronDoNew().UpdatePkData()
		log.Println("修复评课统计数据成功")
	} else {
		var tempPK []global.PkTj
		json.Unmarshal([]byte(pk.Data), &tempPK)
		if len(tempPK) == 0 {
			core.CronDoNew().UpdatePkData()
			log.Println("修复评课统计数据成功")
		}
	}

	log.Println("系统修复自检程序: 修复已完成, 重启评课程序后请前往>>>管理页面<<")
}
