package model

import (
	jsoniter "github.com/json-iterator/go"
	"lzq-admin/domain/model/extrastruct"
	"lzq-admin/pkg/hsflogger"
	"reflect"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/3/29
 * @Version 1.0.0
 */

func (SystemConfig *SystemConfig) TableName() string {
	return TableSystemConfig
}

type SystemConfig struct {
	BaseModel                   `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
	SystemConfigBase            `xorm:"extends"`
	Status                      string `json:"status" xorm:"varchar(32) notnull comment('状态')"` //状态
}

type SystemConfigBase struct {
	ConfigType string `json:"configType" binding:"required" xorm:"varchar(20) notnull comment('配置类型')"` //配置类型
	Code       string `json:"code" binding:"required" xorm:"varchar(32) notnull comment('配置编码')"`       //配置编码
	Name       string `json:"name" binding:"required" xorm:"varchar(32) notnull comment('配置名称')"`       //配置名称
}

type CreateSystemConfigDto struct {
	SystemConfigBase
	ExtraValue interface{} `json:"extraValue" binding:"required"` //扩展字段
}

type UpdateSystemConfigDto struct {
	ID string `json:"id" binding:"required"` //用户id
	SystemConfigBase
	ExtraValue interface{} `json:"extraValue" binding:"required"` //扩展字段
}

type SystemConfigDto struct {
	ID               string `json:"id"` //用户id
	SystemConfigBase `xorm:"extends"`
	ExtraValue       interface{} `json:"extraValue" binding:"required"` //扩展字段
}

type ConfigUpdateDtoBase struct {
	ConfigType string `json:"configType" binding:"required"` //配置类型
	Code       string `json:"code" binding:"required"`       //配置编码
}

type QiNiuConfigDto struct {
	ConfigUpdateDtoBase
	ExtraValue extrastruct.ExtraQiNiuConfig
}

const ExtraSysConfigKey = "ExtraProperty"

const (
	ExtraString      = "ExtraString"
	ExtraInt         = "ExtraInt"
	ExtraStringArray = "ExtraStringArray"
	ExtraIntArray    = "ExtraIntArray"
	ExtraQiNiuConfig = "ExtraQiNiuConfig"
)

var ConfigTypeConstFlags = map[string]string{
	ExtraString:      "字符串",
	ExtraInt:         "整形",
	ExtraStringArray: "字符串数组",
	ExtraIntArray:    "整形数组",
	ExtraQiNiuConfig: "七牛云",
}

var ConfigTypeConstStruct = map[string]interface{}{
	ExtraString:      "",
	ExtraInt:         0,
	ExtraStringArray: make([]string, 0),
	ExtraIntArray:    make([]int, 0),
	ExtraQiNiuConfig: new(extrastruct.ExtraQiNiuConfig),
}

func ReflectSysconfigJsonMap(obj interface{}) map[string]interface{} {
	var tagMap = make(map[string]interface{}, 0)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		hsflogger.LogError("Check type error not Struct", nil)
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		f := t.Field(i)
		tp := f.Type.Kind()
		fieldName := f.Name
		if tp == reflect.Struct {
			vf := v.Field(i)
			v, _ := jsoniter.MarshalToString(vf.Interface())
			mapJson, _ := jsoniter.MarshalToString(reflectStruct(vf.Interface()))
			jM := make(map[string]interface{})
			jM["KeyJson"] = v
			jM["KeyDesc"] = mapJson
			tagMap[fieldName] = jM
		}
	}
	return tagMap
}

func reflectStruct(obj interface{}) map[string]interface{} {
	var tagMap = make(map[string]interface{}, 0)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		hsflogger.LogError("Check type error not Struct", nil)
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		f := t.Field(i)
		tags := f.Tag
		tp := f.Type.Kind()
		fieldName := f.Name
		if len(tags.Get("lab")) > 0 {
			tagMap[fieldName] = tags.Get("lab")
		} else if tp == reflect.Struct {
			vf := v.Field(i)
			tagMap[fieldName] = reflectStruct(vf.Interface())
		}
	}
	return tagMap
}
