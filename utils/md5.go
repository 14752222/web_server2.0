package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成md5
const MD5_LEN = 32

const MD5_STR = "0123456789abcdef"

/*
CreateMd5

	*@param str
	*@return string
	*example:
	*	str := "123456"
	*	md5 := utils.Md5(str)
*/
func CreateMd5(str string) string {
	//创建一个md5对象
	h := md5.New()
	//加密
	h.Write([]byte(str))
	//获取加密结果
	return hex.EncodeToString(h.Sum(nil))
}

/*
EqualMd5

	*@param str1
	*@param str2
	*@return bool
	*example:
	*	str1 := CreateMd5("123456")
	*	str2 := "123456"
	*	equal := utils.EqualMd5(str1, str2)
*/
func EqualMd5(old, str2 string) bool {
	return old == CreateMd5(str2)
}
