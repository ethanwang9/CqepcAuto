package internal

import (
	"CqepcAutoServer/global"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

// name: 日志分割
// author: Ethan.Wang
// desc: 使用 file-rotatelogs 分割日志, 因为该库已只读，所以使用了 gopkg.in/natefinch/lumberjack.v2

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	t := time.Now().Format("2006-01-02")
	lumberJackLogger := &lumberjack.Logger{
		// 日志文件的位置
		Filename: path.Join(global.CONFIG.GetString("log.dir"), t, level+".log"),
		// 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxSize: global.CONFIG.GetInt("log.max_size"),
		// 保留旧文件的最大个数
		MaxBackups: global.CONFIG.GetInt("log.max_backups"),
		// 保留旧文件的最大天数
		MaxAge: global.CONFIG.GetInt("log.max_age"),
		// 是否压缩/归档旧文件
		Compress: true,
	}

	if global.CONFIG.GetBool("log.log_in_console") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
