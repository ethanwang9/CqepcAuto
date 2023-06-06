package utils

import (
	"encoding/base64"
	"github.com/axelwong/CqepcAuto/global"
	"go.uber.org/zap"
)

// name: 算法
// author: axel wong
// desc: 常用算法

type Algorithm struct{}

var AlgorithmApp = new(Algorithm)

// Base64Encode Base64 编码
func (a *Algorithm) Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64Decode Base64 解码
func (a *Algorithm) Base64Decode(dst string) (src []byte, err error) {
	src, err = base64.StdEncoding.DecodeString(dst)
	if err != nil {
		global.APP_LOG.Warn(
			"工具类-base64解码失败",
			zap.String("dst", dst),
			zap.String("error", err.Error()),
		)
		return nil, err
	}
	return src, nil
}
