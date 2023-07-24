package utils

import (
	"CqepcAuto/global"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

// GetStringHash256 hash256
func (a *Algorithm) GetStringHash256(text string) string {
	//1.new一个指定的hash函数
	stringHash := sha256.New() //返回hash.Hash
	//2.向hash中添加数据
	stringHash.Write([]byte(text))
	//3. 计算hash结果
	temp := stringHash.Sum(nil)
	//4. 固定长度的字符串
	hash := hex.EncodeToString(temp)
	return hash //返回哈希
}
