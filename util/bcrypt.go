package util

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成加密字符串
func BcryptString(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// 验证用户输入的和加密后的是否一致
func VerifyString(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	if err != nil {
		return false
	}

	return true
}
