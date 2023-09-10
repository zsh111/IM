package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//加密模块，实现对用户的密码进行加密和解密

// 转小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempstr := h.Sum(nil)
	return hex.EncodeToString(tempstr)
}

// 转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 将传入数据加随机数一同加密
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
