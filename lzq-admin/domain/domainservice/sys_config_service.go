package domainservice

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/domain/model/extrastruct"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/3/29
 * @Version 1.0.0
 */

type sysConfigDomainService struct{}

var SysConfigDomainService = sysConfigDomainService{}

func (s *sysConfigDomainService) Insert(inputDto model.CreateSystemConfigDto) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(true).Where("ConfigType=? and Code=?", inputDto.ConfigType, inputDto.Code).Exist(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if isExit {
		return systemConfig, errors.New("该配置类型已存在该配置编码：" + inputDto.Code)
	}

	systemConfig.ID = utility.UuidCreate()
	systemConfig.ConfigType = inputDto.ConfigType
	systemConfig.Code = inputDto.Code
	systemConfig.Name = inputDto.Name
	systemConfig.ExtraProperties = make(map[string]interface{})
	systemConfig.ExtraProperties[model.ExtraSysConfigKey] = model.ConfigTypeConstStruct[inputDto.ConfigType] //s.GetExtraPropertiesJson(inputDto.ConfigType) //new(extrastruct.ExtraQiNiuConfig)
	systemConfig.Status = domainconsts.CommonStatusEnable
	if _, err := orm.ISession().Insert(&systemConfig); err != nil {
		return model.SystemConfig{}, err
	}
	return systemConfig, nil
}

func (s *sysConfigDomainService) Update(inputDto model.UpdateSystemConfigDto) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(false).Where("ConfigType=? and Code=?", inputDto.ConfigType, inputDto.Code).Get(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if !isExit {
		return systemConfig, errors.New("该配置类型不已存在：" + inputDto.Code)
	}
	systemConfig.ExtraProperties[model.ExtraSysConfigKey] = inputDto.ExtraValue
	if _, err := orm.USession(true).AllCols().ID(systemConfig.ID).Update(&systemConfig); err != nil {
		return model.SystemConfig{}, err
	}
	r := cache.RedisUtil.NewRedis(false, "SysConfig")
	r.Delete(fmt.Sprintf("%v:%v", systemConfig.ConfigType, systemConfig.Code))
	return systemConfig, nil
}

func (s *sysConfigDomainService) GetByCode(configType, code string) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(false).Where("ConfigType=? and Code=?", configType, code).Get(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if !isExit {
		return systemConfig, errors.New("不存在该配置")
	}

	return systemConfig, nil
}

func (s *sysConfigDomainService) GetSysConfigCacheByCode(configType, code string) (string, error) {
	r := cache.RedisUtil.NewRedis(false, "SysConfig")
	key := fmt.Sprintf("%v:%v", configType, code)
	obj := r.Get(key)
	if obj != "" {
		return obj, nil
	}
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(false).Where("ConfigType=? and Code=?", configType, code).Get(&systemConfig)
	if err != nil {
		return "", err
	}
	if !isExit {
		return "", errors.New("不存在该配置")
	}
	v := systemConfig.ExtraProperties[model.ExtraSysConfigKey]
	r.SSet(key, v, 0)
	json, _ := jsoniter.MarshalToString(v)
	return json, nil
}

func (s *sysConfigDomainService) GetExtraPropertiesJson(configType string) interface{} {
	var json interface{}
	switch configType {
	case model.ExtraString:
		json = ""
		break
	case model.ExtraInt:
		json = 0
		break
	case model.ExtraStringArray:
		json = make([]string, 0)
		break
	case model.ExtraIntArray:
		json = make([]int, 0)
		break
	case model.ExtraQiNiuConfig:
		json = new(extrastruct.ExtraQiNiuConfig)
		break
	default:
		json = s.SetSysConfigJsonMapCache(configType)
		break
	}
	return json
}

func (s *sysConfigDomainService) SetSysConfigJsonMapCache(configType string) interface{} {
	c := cache.RedisUtil.NewRedis(false)
	var objJson interface{}
	isAllConfigType := configType == "AllConfigType"
	if !isAllConfigType {
		objJson = c.HGet("ExtraStruct:KeyJson", configType)
		if objJson != "" {
			return objJson
		}
	}
	var obj extrastruct.ExtraConfigObject
	jsonMap := model.ReflectSysconfigJsonMap(obj)
	if isAllConfigType {
		for k, v := range jsonMap {
			vm := v.(map[string]interface{})
			for kk, vv := range vm {
				c.HSet(fmt.Sprintf("ExtraStruct:%v", kk), k, vv)
			}
		}
		objJson = jsonMap
	} else {
		if v, ok := jsonMap[configType]; !ok {
			errors.New("该配置类型：" + configType + " 不存在")
		} else {
			vm := v.(map[string]interface{})
			objJson = vm["KeyJson"]
		}
	}

	return objJson
}
