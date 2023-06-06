package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

type Tool struct{}

var ToolApp = new(Tool)

// RandString 随机字符串
// mode 1-小写, 2-大写, 其他-不大小写转义
func (t *Tool) RandString(length, mode int) string {
	rand.Seed(time.Now().UnixNano())
	randCore := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	randStr := ""
	for i := 0; i < length; i++ {
		randStr += string(randCore[rand.Intn(len(randCore))])
	}
	switch mode {
	case 1:
		// 小写
		randStr = strings.ToLower(randStr)
	case 2:
		// 大写
		randStr = strings.ToUpper(randStr)
	default:
		break
	}
	return randStr
}

// RandWxOpenid 生成随机微信openid
func (t *Tool) RandWxOpenid() (s string) {
	rand.Seed(time.Now().UnixNano())
	// 初始化字符串
	s = "oI2Hl5"
	// 是否间隔
	if ge := rand.Intn(10); ge%2 == 1 {
		// 输出间隔
		b := t.RandString(rand.Intn(15)+6, 0)
		s += b + "-" + t.RandString(21-len(b), 0)
	} else {
		// 无间隔
		s += t.RandString(22, 0)
	}
	return
}

// UnixToDate 时间戳 -> 日期
func (t *Tool) UnixToDate(timeStamp int64) string {
	timeFormat := "2006-01-02 15:04"
	Loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		Loc = time.FixedZone("CST", 8*3600)
	}

	t3 := time.Unix(timeStamp, 0).In(Loc)
	return t3.Format(timeFormat)
}

// DateToUnix 日期 -> 时间戳
func (t *Tool) DateToUnix(d string) int64 {
	timeFormat := "2006-01-02 15:04"
	Loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		Loc = time.FixedZone("CST", 8*3600)
	}

	t2, _ := time.ParseInLocation(timeFormat, d, Loc)
	return t2.Unix()
}

// UUID 生成UUID
func (t *Tool) UUID() (u string) {
	u = uuid.NewV4().String()
	u = strings.Replace(u, "-", "", -1)
	return u
}

// TodayZeroUnix 当天 00:00:00 时间戳
func (t *Tool) TodayZeroUnix() time.Time {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}

	date := fmt.Sprintf("%v-%v-%v 00:00:00", time.Now().Year(), time.Now().Format("01"), time.Now().Day())
	tUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)

	return tUnix
}

// TodayFullUnix 当天 23:59:59 时间戳
func (t *Tool) TodayFullUnix() time.Time {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}

	date := fmt.Sprintf("%v-%v-%v 23:59:59", time.Now().Year(), time.Now().Format("01"), time.Now().Day())
	tUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)

	return tUnix
}

// TodayOneUnix 当天 01:00:00 时间戳
func (t *Tool) TodayOneUnix() time.Time {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}

	date := fmt.Sprintf("%v-%v-%v 1:00:00", time.Now().Year(), time.Now().Format("01"), time.Now().Day())
	tUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)

	return tUnix
}
