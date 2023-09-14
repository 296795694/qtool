package entitystruct

import (
	"fmt"
	"reflect"
)

// StructToMapReflect 结构体转map
func StructToMapReflect(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("结构体格式不正确 %T", v)
	}
	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := GetReflectTag(fi.Type); tagValue != "" {
			out[tagValue.Get("json")] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// GetReflectTag 无限极获取tag标签
func GetReflectTag(reflectType reflect.Type) (tag reflect.StructTag) {
	if reflectType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < reflectType.NumField(); i++ {
		tag = reflectType.Field(i).Tag
		if tag == "" {
			tag = GetReflectTag(reflectType.Field(i).Type)
			continue
		}
	}
	return
}
