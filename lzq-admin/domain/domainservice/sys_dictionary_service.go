package domainservice

import (
	"errors"
	"fmt"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/7/5
 * @Version 1.0.0
 */

type systemDictionaryDomainService struct {
}

var SystemDictionaryDomainService = systemDictionaryDomainService{}

func (s *systemDictionaryDomainService) Insert(inputDto model.CreateSystemDictionaryDto) (model.SystemDictionary, error) {
	var entity model.SystemDictionary

	if isExist, err := orm.QSession(false).Where("DictCode = ? and DictKey = ?", inputDto.DictCode, inputDto.DictKey).Exist(&entity); err != nil {
		return entity, err
	} else if isExist {
		return entity, errors.New(fmt.Sprintf("该字典Key：%s已存在，请更换字典Key", inputDto.DictKey))
	}

	entity.ID = utility.UuidCreate()
	entity.DictCode = inputDto.DictCode
	entity.DictKey = inputDto.DictKey
	entity.DictValue = inputDto.DictValue
	entity.Remark = inputDto.Remark
	entity.Sort = inputDto.Sort
	entity.Status = domainconsts.CommonStatusEnable
	if _, err := orm.ISession().Insert(&entity); err != nil {
		return entity, err
	}

	cacheKey := cache.LzqCacheKeyHelper.GetLzqCacheTypeSystemDict(entity.DictCode)
	cacheObj := s.SystemDictToSimple(entity)
	fmt.Println(cacheKey, cacheObj)
	cache.RedisUtil.NewRedis(false).HSet(cacheKey, entity.DictKey, cacheObj, 0)
	return entity, nil
}

func (s *systemDictionaryDomainService) Update(inputDto model.UpdateSystemDictionaryDto) (model.SystemDictionary, error) {
	var entity model.SystemDictionary
	if isExist, err := orm.QSession(false).ID(inputDto.ID).Get(&entity); err != nil {
		return entity, err
	} else if !isExist {
		return entity, errors.New("字典不存在")
	}

	entity.DictValue = inputDto.DictValue
	entity.Remark = inputDto.Remark
	entity.Sort = inputDto.Sort
	if _, err := orm.USession(false).AllCols().ID(entity.ID).Update(&entity); err != nil {
		return entity, err
	}
	s.UpdateSystemDictCache(entity)
	return entity, nil
}

func (s *systemDictionaryDomainService) UpdateStatus(id string) error {
	var entity model.SystemDictionary
	if isExist, err := orm.QSession(false).ID(id).Get(&entity); err != nil {
		return err
	} else if !isExist {
		return errors.New("字典不存在")
	}
	if entity.Status == domainconsts.CommonStatusEnable {
		entity.Status = domainconsts.CommonStatusDisable
	} else {
		entity.Status = domainconsts.CommonStatusEnable
	}
	if _, err := orm.USession(false).Cols("Status").ID(id).Update(&entity); err != nil {
		return err
	}
	s.UpdateSystemDictCache(entity)
	return nil
}

func (s *systemDictionaryDomainService) Delete(id string) error {
	var entity model.SystemDictionary
	if isExist, err := orm.QSession(false).ID(id).Get(&entity); err != nil {
		return err
	} else if !isExist {
		return errors.New("字典不存在")
	}

	if _, err := orm.DSession(false).ID(id).Update(&entity); err != nil {
		return err
	}

	cacheKey := cache.LzqCacheKeyHelper.GetLzqCacheTypeSystemDict(entity.DictCode)
	cache.RedisUtil.NewRedis(false).HDelete(cacheKey, entity.DictKey)
	return nil
}

func (s *systemDictionaryDomainService) GetDictByDictKey(dictCode, dictKey string) (model.SimpleSystemDictionaryDto, error) {
	var entity model.SystemDictionary
	var result model.SimpleSystemDictionaryDto

	cacheKey := cache.LzqCacheKeyHelper.GetLzqCacheTypeSystemDict(dictCode)
	cacheJson := cache.RedisUtil.NewRedis(false).HGet(cacheKey, dictKey)
	if cacheJson != nil {
		s, ok := cacheJson.(model.SimpleSystemDictionaryDto)
		if ok {
			return s, nil
		}
	}

	if isExist, err := orm.QSession(false).Where("DictCOde = ? and DictKey = ?", dictCode, dictKey).Get(&entity); err != nil {
		return result, err
	} else if !isExist {
		return result, errors.New(fmt.Sprintf("该字典：%s下不存在该Key：%s", dictCode, dictKey))
	}
	result = s.SystemDictToSimple(entity)
	cache.RedisUtil.NewRedis(false).HSet(cacheKey, dictKey, result, 0)

	return result, nil
}

func (s *systemDictionaryDomainService) GetDictByDictCode(dictCode string) ([]model.SimpleSystemDictionaryDto, error) {
	var list = make([]model.SystemDictionary, 0)
	var result = make([]model.SimpleSystemDictionaryDto, 0)

	cacheKey := cache.LzqCacheKeyHelper.GetLzqCacheTypeSystemDict(dictCode)
	cacheJson := cache.RedisUtil.NewRedis(false).HGetAll(cacheKey)
	if cacheJson != "" {
		s, ok := cacheJson.([]model.SimpleSystemDictionaryDto)
		if ok {
			return s, nil
		}
	}

	if err := orm.QSession(false).Where("DictCOde = ?", dictCode).Find(&list); err != nil {
		return result, err
	}
	for _, v := range list {
		m := s.SystemDictToSimple(v)
		result = append(result, m)
		cache.RedisUtil.NewRedis(true).HSet(cacheKey, v.DictKey, m, 0)
	}
	return result, nil
}

func (s *systemDictionaryDomainService) SystemDictToSimple(v model.SystemDictionary) model.SimpleSystemDictionaryDto {
	return model.SimpleSystemDictionaryDto{
		SystemDictionaryBase: model.SystemDictionaryBase{
			DictCode:  v.DictCode,
			DictKey:   v.DictKey,
			DictValue: v.DictValue,
			Sort:      v.Sort,
			Remark:    v.Remark,
		},
		Status: v.Status,
	}
}

func (s *systemDictionaryDomainService) UpdateSystemDictCache(v model.SystemDictionary) {
	cacheKey := cache.LzqCacheKeyHelper.GetLzqCacheTypeSystemDict(v.DictCode)
	cacheObj := cache.RedisUtil.NewRedis(false).HGet(cacheKey, v.DictKey)
	if cacheObj != nil {
		cache.RedisUtil.NewRedis(false).HDelete(cacheKey, v.DictKey)
	}
	cacheObj = s.SystemDictToSimple(v)
	cache.RedisUtil.NewRedis(false).HSet(cacheKey, v.DictKey, cacheObj, 0)
}

func(s *systemDictionaryDomainService) RefreshSystemDictCache() error{
	cache.RedisUtil.NewRedis(false).Delete(cache.LzqCacheHelper.GetCacheTypeVersionKey(cache.LzqCacheTypeSystemDict))
	return nil
}
