package controller

import (
	"CqepcAuto/global"
	"CqepcAuto/model"
	"CqepcAuto/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AdminIndex(ctx *gin.Context) {
	// 获取用户信息
	userInfo, err := model.UserNew(model.User{}).GetByUidOldest()
	if err != nil {
		ctx.String(http.StatusOK, "获取数据库@用户信息异常，请再次重试！")
		return
	}

	// 获取今日课表
	dbToday, err := model.TodayNew(model.Today{Uid: userInfo.Uid}).GetByUidToday()
	if err != nil {
		ctx.String(http.StatusOK, "获取数据库@今日课表信息异常，请再次重试！")
		return
	}
	today := make([]map[string]interface{}, 0)
	var todayTime time.Time
	for _, v := range dbToday {
		var tempData global.ResTodayClassData
		json.Unmarshal([]byte(v.Data), &tempData)

		isPk := ""
		switch v.IsPk {
		case "no":
			isPk = "未评课"
		case "success":
			isPk = "已评课"
		case "danger":
			isPk = "已补评"
		}

		if v.CreatedAt.Unix() >= utils.ToolApp.TodayZeroUnix().Unix() && v.CreatedAt.Unix() <= utils.ToolApp.TodayOneUnix().Unix() {
			isPk = "课表已过期，等待更新中..."
		}

		today = append(today, gin.H{
			"className":   tempData.CourseName,
			"classRoom":   tempData.ClassClassromm,
			"teacherName": tempData.TeacherName,
			"isPk":        isPk,
			"startTime":   tempData.NodeBegin,
			"endTime":     tempData.NodeEnd,
			"classCode":   v.PId,
		})

		todayTime = v.CreatedAt
	}

	// 获取评课信息
	dbTj, err := model.TjNew(model.Tj{Uid: userInfo.Uid}).GetByUid()
	if err != nil {
		ctx.String(http.StatusOK, "获取数据库@获取评课信息异常，请再次重试！")
		return
	}

	var tempDB []global.PkTj
	json.Unmarshal([]byte(dbTj.Data), &tempDB)
	if len(tempDB) == 0 {
		ctx.String(http.StatusOK, "数据库无评课数据，请运行系统修复自检程序")
		return
	}

	tj := make([]map[string]interface{}, 0)
	for _, v := range tempDB {
		tj = append(tj, gin.H{
			"courseName":      v.CourseName,
			"courseCode":      v.CourseCode,
			"totalClassroom":  v.Data.TotalClassroom,
			"totalEvaluation": v.Data.TotalEvaluation,
			"prob":            fmt.Sprintf("%3.2f%%", float32(v.Data.TotalEvaluation)/float32(v.Data.TotalClassroom)*100),
			"startTime":       v.Data.StartTime,
			"endTime":         v.Data.EndTime,
		})
	}

	// 错误修正
	if todayTime.Unix() < 0 {
		todayTime = time.Now()
	}

	// 返回数据
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"userinfo": userInfo,
		"app": gin.H{
			"now":         "Version " + global.Version,
			"author":      global.Author,
			"release_url": global.AppUrl,
		},
		"today": today,
		"tj":    tj,
		"now":   time.Now().Unix(),
		"time": gin.H{
			"today": todayTime.Format("2006-01-02 15:04:05"),
			"tj":    dbTj.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		"count": gin.H{
			"today": len(today),
		},
	})
}
