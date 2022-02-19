package utility

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"encoding/json"
	"reflect"
)

// StructToMap struct转为map
func StructToMap(obj interface{}, oneSeries bool) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		if oneSeries == true {
			if v.Field(i).Kind() == reflect.Struct {
				m := StructToMap(v.Field(i).Interface(), false)
				//fmt.Println(m)
				for k, w := range m {
					//fmt.Println(k, w)
					data[k] = w
				}
			} else {
				data[fieldName] = v.Field(i).Interface()
			}
		} else {
			data[fieldName] = v.Field(i).Interface()
		}
	}

	return data
}

// StructToJson struct转为json
//这里对应的 N 和 A 不能为小写，首字母必须为大写，这样才可对外提供访问，具体 json 匹配是通过后面的 tag 标签进行匹配的，与 N 和 A 没有关系
//tag 标签中 json 后面跟着的是字段名称，都是字符串类型，要求必须加上双引号，否则 golang 是无法识别它的类型
func StructToJson(obj interface{}) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
