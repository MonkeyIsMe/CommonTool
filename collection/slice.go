package collection

import (
	"math"

	"github.com/MonkeyIsMe/MyTool/constant"
)

// GetListByPage 字符串数组分页
func GetListByPage(talentType int32, prevPage int32, list []string) ([]string, bool) {
	totalLen := len(list)                                                     // 总共数据长度
	maxPage := int(math.Ceil(float64(totalLen) / float64(constant.PageSize))) // 总共的数据页数
	page := (prevPage - 1) * int32(constant.PageSize)                         // 第n页的开头
	limit := prevPage * int32(constant.PageSize)                              // 第n页的结尾

	// 如果数据总长度小于每页长度直接返回
	if constant.PageSize > totalLen {
		ansList := list[0:totalLen]
		return ansList, true
	}

	// 如果请求的页数超过范围
	if prevPage > int32(maxPage) {
		return nil, true
	}

	// 如果本次请求的数据到了最后
	if limit > int32(totalLen) {
		ansList := list[page:totalLen]
		return ansList, true
	}

	// 正常分页
	ansList := list[page:limit]

	return ansList, false
}
