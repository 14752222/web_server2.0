package utils

import (
	"math/rand"
	"time"
)

func CreateCheckCode(length int, include bool) string {
	//生成随机数
	//length:生成随机数的长度
	//include:是否包含大小写字母
	//返回值:生成的随机数
	if length <= 0 {
		length = 6
	}

	var code string

	if include {
		code = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		code = "0123456789"
	}

	var result []byte
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, code[randSource.Intn(len(code))])
	}

	return string(result)
}
