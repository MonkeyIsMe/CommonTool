package strconv

import (
	"log"
	"strconv"
	"strings"
)

// ParseIntList 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseIntList(s string, sep string) []int {
	ret64 := ParseInt64List(s, sep)
	ret := make([]int, 0, len(ret64))
	for _, v := range ret64 {
		ret = append(ret, int(v))
	}
	return ret
}

// ParseInt32List 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseInt32List(s string, sep string) []int32 {
	ret64 := ParseInt64List(s, sep)
	ret := make([]int32, 0, len(ret64))
	for _, v := range ret64 {
		ret = append(ret, int32(v))
	}
	return ret
}

// ParseInt64List 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseInt64List(s string, sep string) []int64 {
	l := strings.Split(s, sep)
	ret := make([]int64, 0, len(l))
	for _, v := range l {
		v := strings.TrimSpace(v)
		if v == "" {
			continue
		}

		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatalf("strconv.ParseInt error, %+v, v:%s", err, v)
			continue
		}

		ret = append(ret, i)
	}
	return ret
}

// ParseUintList 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseUintList(s string, sep string) []uint {
	ret64 := ParseUint64List(s, sep)
	ret := make([]uint, 0, len(ret64))
	for _, v := range ret64 {
		ret = append(ret, uint(v))
	}
	return ret
}

// ParseUint32List 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseUint32List(s string, sep string) []uint32 {
	ret64 := ParseUint64List(s, sep)
	ret := make([]uint32, 0, len(ret64))
	for _, v := range ret64 {
		ret = append(ret, uint32(v))
	}
	return ret
}

// ParseUint64List 将s用sep分开，然后逐个转换为整数。转换失败的内容，直接跳过。
func ParseUint64List(s string, sep string) []uint64 {
	l := strings.Split(s, sep)
	ret := make([]uint64, 0, len(l))
	for _, v := range l {
		v := strings.TrimSpace(v)
		if v == "" {
			continue
		}

		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Fatalf("strconv.ParseUint error, %+v, v:%s", err, v)
			continue
		}

		ret = append(ret, i)
	}
	return ret
}
