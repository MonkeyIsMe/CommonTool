package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"time"
)

// VerifyEmailFormat 验证邮箱合法性
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobileFormat 验证手机号合法性
func VerifyMobileFormat(mobileNumber string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNumber)
}

// SignIDCard 身份证号码处理
func SignIDCard(idcard string) string {
	cp := idcard
	leth := len(cp)
	return cp[0:4] + " **** **** " + cp[leth-4:]
}

// SignMobilePhone 手机号码处理
func SignMobilePhone(mobile string) string {
	cp := mobile
	leth := len(cp)
	return cp[0:3] + " **** " + cp[leth-4:]
}

// GetUniqueId 生成UniqueId方法
func GetUniqueId() string {
	// 当前毫秒时间戳
	timestamp := time.Now().UnixNano() / 1000000
	s1, _ := rand.Int(rand.Reader, big.NewInt(10))
	s2, _ := rand.Int(rand.Reader, big.NewInt(10))

	seqNo := timestamp * (s1.Int64() + 1) * (s2.Int64() + 1)
	uniqueId := fmt.Sprintf("%d", seqNo)
	return uniqueId
}

// VerifyPassword 验证密码合法性
func VerifyPassword(password string) bool {
	len := len(password)
	digitFlag := 0
	bigLetterFlag := 0
	smallLetterFlag := 0
	specialLetterFlag := 0
	for i := 0; i < len; i++ {
		if password[i] >= '0' && password[i] <= '9' {
			digitFlag = 1
		} else if password[i] >= 'a' && password[i] <= 'z' {
			smallLetterFlag = 1
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			bigLetterFlag = 1
		} else if password[i] == '!' || password[i] == '~' || password[i] == '@' || password[i] == '#' ||
			password[i] == '_' || password[i] == '%' || password[i] == '^' || password[i] == '*' {
			specialLetterFlag = 1
		} else {
			return false
		}
	}
	flag := digitFlag + smallLetterFlag + specialLetterFlag + bigLetterFlag
	if len >= 6 && len <= 15 && flag >= 3 {
		return true
	}

	return false
}
