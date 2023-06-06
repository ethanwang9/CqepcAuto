package core

import (
	"fmt"
	"github.com/axelwong/CqepcAuto/global"
	"github.com/axelwong/CqepcAuto/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// name: 日志 zap
// author: axel wong
// desc: 处理日志文件

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.Director); !ok {
		_ = os.Mkdir(global.Director, os.ModePerm)
	}
	// 调试级别: 可以随意的使用于任何觉得有利于在调试时更详细的了解系统运行状态
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别: 用来反馈系统的当前状态给最终用户的
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别: 不影响运行的错误，一般为网络错误，请求接口失败等
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别: 影响运行的错误，一般为数据库连接失败，定时器未生效等
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	})
	// 异常级别: 错误已经无法修复，系统会停止运行
	fatalPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.DPanicLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.Director), errorPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_fatal.log", global.Director), fatalPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch global.EncodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	// 使用 file-rotatelogs 进行日志分割
	writer := utils.GetWriteSyncer(fileName)
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Prefix + "2006-01-02 15:04:05.000"))
}
