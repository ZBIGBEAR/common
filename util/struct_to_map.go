package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
)

/* StructToMapV1.
结构体转换成map，空字段会忽略。使用序列化和反序列化方法。
*/
func StructToMapV1(src interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return nil, errors.Wrapf(err, "[StructToMapV1] Marshal src:%+v", src)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Wrapf(err, "[StructToMapV1] Unmarshal src:%+v", src)
	}

	return result, nil
}

/* StructToMap.
结构体转换成map,空字段不会忽略。使用反射方法
*/
func StructToMapV2(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	tagName := "json"
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			tagValues := strings.Split(tagValue, ",")
			fieldName := strings.Trim(tagValues[0], " ")
			out[fieldName] = v.Field(i).Interface()
		}
	}
	return out, nil
}

/* StructToMapV3.
结构体转换成map.内部也是使用的反射方法，需要指定 tag "structs"
*/
func StructToMapV3(in interface{}) map[string]interface{} {
	return structs.Map(in)
}
