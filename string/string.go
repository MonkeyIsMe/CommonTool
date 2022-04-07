package string

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

const (
	NotFound = -1
)

// SplitString 返回以partition分割的字符串数组
func SplitString(str, partition string) []string {
	splitStr := strings.Split(str, partition)
	return splitStr
}

// IndexItemInSlice 返回item在l中的位置，如果不存在返回-1。
func IndexItemInSlice(l []string, item string) int {
	for pos, v := range l {
		if v == item {
			return pos
		}
	}
	return NotFound
}

// CaclMD5 返回str内容对应md5值(16进制表示)
func CaclMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// FilterEmptyString 过滤空字符串
func FilterEmptyString(src []string) []string {
	dest := make([]string, 0, len(src))
	for _, s := range src {
		if s == "" {
			continue
		}

		dest = append(dest, s)
	}
	return dest
}

// TrimSpaceAndFilterEmpty 将src中的元素去掉头尾空白字符，并丢弃空字符串。多用于配置文件处理。
func TrimSpaceAndFilterEmpty(src []string) []string {
	dest := make([]string, 0, len(src))
	for _, s := range src {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		dest = append(dest, s)
	}
	return dest
}

// GetSubstring 截取字符串 start 起点下标 length 需要截取的长度
func GetSubstring(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// Base64Encode 对数据进行 base64 编码
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode 对数据进行 base64 解码
func Base64Decode(s string) (string, error) {
	rs, err := base64.StdEncoding.DecodeString(s)
	return string(rs), err
}

// GetSha 对字符串进行sha1 计算
func GetSha(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// GetMd5 对数据进行md5计算
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetRandomString 随机生成给定长度的字符串
func GetRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetMD5 生成 md5 数字串
func GetMD5(raw string) uint64 {
	h := md5.New()
	_, _ = io.WriteString(h, raw)
	md5Val := binary.LittleEndian.Uint64(h.Sum(nil)[0:8])
	return md5Val
}

// 生成 md5 哈希
func GetHexMD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// string转[]byte无拷贝
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// InterfaceToString 接口转string
func InterfaceToString(inter interface{}) (string, error) {
	switch inter.(type) {
	case string:
		return inter.(string), nil
	default:
		return "", fmt.Errorf("Type Error")
	}
}

// Json2String json对象转换为字符串
func Json2String(d interface{}) (string, error) {
	j, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
