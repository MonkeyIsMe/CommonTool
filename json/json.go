package json

import "encoding/json"

// String2Json 字符串转换为json
func String2Json(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}

// Map2JsonString map对象转化成json string
func Map2JsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
