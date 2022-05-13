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
	systemConfig.ExtraProperties[model.ExtraSysConfigKey] = inputDto.ExtraValue
	systemConfig.Status = domainconsts.CommonStatusEnable
	if _, err := orm.ISession().Insert(&systemConfig); err != nil {
		return model.SystemConfig{}, err
	}

	return systemConfig, nil
}

func (s *sysConfigDomainService) Update(inputDto model.UpdateSystemConfigDto) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(true).Where("ConfigType=? and Code=?", inputDto.ConfigType, inputDto.Code).Get(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if !isExit {
		return systemConfig, errors.New("该配置类型不已存在：" + inputDto.Code)
	}

	systemConfig.Name = inputDto.Name
	systemConfig.ExtraProperties[model.ExtraSysConfigKey] = inputDto.ExtraValue
	systemConfig.Status = domainconsts.CommonStatusEnable
	if _, err := orm.USession(true).ID(inputDto.ID).Update(&systemConfig); err != nil {
		return model.SystemConfig{}, err
	}

	return systemConfig, nil
}

func (s *sysConfigDomainService) GetById(id string) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(true).ID(id).Get(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if !isExit {
		return systemConfig, errors.New("不存在该配置")
	}
	return systemConfig, nil
}

func (s *sysConfigDomainService) GetByCode(configType, code string) (model.SystemConfig, error) {
	var systemConfig model.SystemConfig
	isExit, err := orm.QSession(true).Where("ConfigType=? and Code=?", configType, code).Get(&systemConfig)
	if err != nil {
		return systemConfig, err
	}
	if !isExit {
		return systemConfig, errors.New("不存在该配置")
	}
	return systemConfig, nil
}

func (s *sysConfigDomainService) GetExtraPropertiesJson(configType string) interface{} {
	var json interface{}
	switch configType {
	case model.ExtraString:
		break
	case model.ExtraInt:
		json = 0
		break
	case model.ExtraStringArray:
		var t []string
		json, _ = jsoniter.MarshalToString(t)
		break
	case model.ExtraIntArray:
		var t []int
		json, _ = jsoniter.MarshalToString(t)
		break
	default:
		json = s.SetSysconfigJsonMapCache(configType)
		break
	}
	return json
}

func (s *sysConfigDomainService) SetSysconfigJsonMapCache(configType string) interface{} {
	c := cache.RedisUtil.NewRedis(false)
	var objJson interface{}
	isAllConfigType := configType == "AllConfigType"
	if !isAllConfigType {
		objJson = c.HGet("Extrastruct:KeyJson", configType)
		if objJson != nil {
			return objJson
		}
	}
	var obj extrastruct.ExtraConfigObject
	jsonMap := model.ReflectSysconfigJsonMap(obj)
	if isAllConfigType {
		for k, v := range jsonMap {
			vm := v.(map[string]interface{})
			for km, vm := range vm {
				c.HSet(fmt.Sprintf("Extrastruct:%v", km), k, vm)
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
