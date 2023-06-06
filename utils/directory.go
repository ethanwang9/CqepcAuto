package utils

import (
	"github.com/axelwong/CqepcAuto/global"
	"go.uber.org/zap"
	"os"
)

// name: IO工具类
// author: axel wong
// desc:

// PathExists 判断 文件|文件夹 是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.APP_LOG.Debug("工具类-创建文件夹" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.APP_LOG.Error("工具类-创建文件夹失败"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
