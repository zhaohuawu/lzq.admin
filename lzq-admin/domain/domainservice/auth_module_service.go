package domainservice

import (
	"github.com/pkg/errors"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/1
 * @Version 1.0.0
 */

// 服务模块领域服务
type authModuleDomainService struct {
	wg sync.WaitGroup
}

var AuthModuleDomainService = authModuleDomainService{}

func (*authModuleDomainService) Insert(modelDto model.CreateAuthModuleDto) (result model.AuthModule, err error) {
	var entity model.AuthModule
	var exist bool
	exist, err = orm.QSession(false).Where("code=?", modelDto.Code).Exist(&model.AuthModule{})
	if err != nil {
		return entity, err
	}
	if exist {
		return entity, errors.New("模块编码：" + modelDto.Code + " 已存在，请更换模块编码")
	}
	entity = model.AuthModule{
		AuthModuleBase: model.AuthModuleBase{
			Code: modelDto.Code,
			Name: modelDto.Name,
			Rank: modelDto.Rank,
		},
	}
	entity.ID = utility.UuidCreate()
	_, err = orm.ISession().Insert(&entity)
	return entity, err
}

func (*authModuleDomainService) Get(id, code string) (model.AuthModuleDto, error) {
	var has bool = false
	var err error
	var dto model.AuthModuleDto
	db := orm.QSession(false)
	if len(id) > 0 {
		has, err = db.ID(id).Get(&dto)
	} else if len(code) > 0 {
		has, err = db.Where("Code=?", code).Get(dto)
	}
	if err != nil {
		return dto, err
	}
	if has {
		return dto, nil
	} else {
		return dto, errors.New("模块ID/模块编码不存在或已删除")
	}
}
