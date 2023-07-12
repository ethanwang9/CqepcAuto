package utils

import (
	"CqepcAutoServer/global"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
)

// name:文件目录是否存在
// author: Ethan.Wang, GVA [github.com/flipped-aurora/gin-vue-admin]
// desc: 关于文件的相关操作

// PathExists 判断文件是否存在
// value: true-存在 | false-不存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return true, errors.New("存在同名文件")
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
			global.LOG.Debug("创建文件夹：" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.LOG.Error("创建文件夹失败",
					zap.String("DirName", v),
					zap.Any("error", err),
				)
				return err
			}
		}
	}
	return err
}

// CreateFile 创建文件并写入内容
func CreateFile(path, name, content string) (err error) {
	file, err := os.Create(fmt.Sprintf("%v%v.json", path, name))
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		global.LOG.Warn("创建文件并写入内容失败#创建文件",
			zap.Error(err),
			zap.Any("data", map[string]interface{}{
				"path": path,
				"name": name,
				"data": content,
			}),
		)
		return err
	}

	// 写入内容
	_, err = file.WriteString(content)
	if err != nil {
		global.LOG.Warn("创建文件并写入内容失败#写入内容",
			zap.Error(err),
			zap.Any("data", map[string]interface{}{
				"path": path,
				"name": name,
				"data": content,
			}),
		)
		return err
	}

	return nil
}
