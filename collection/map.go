package collection

import (
	"reflect"
)

// InsertStringMap 插入map的简单封装，主要是判断map是否是空的，如果为空就构造一个map
func InsertStringMap(srcMap map[string]string, key, value string) map[string]string {
	if srcMap == nil {
		srcMap = map[string]string{
			key: value,
		}
		return srcMap
	}
	srcMap[key] = value
	return srcMap
}

// StructToMap struct转为map 因为struct是不能遍历的
func StructToMap(stu interface{}) map[string]interface{} {
	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}
