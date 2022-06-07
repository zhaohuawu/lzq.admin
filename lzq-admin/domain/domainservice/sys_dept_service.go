package domainservice

import (
	"errors"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/23
 * @Version 1.0.0
 */

type systemDeptDomainService struct {
}

var SystemDeptDomainService = systemDeptDomainService{}

func (s *systemDeptDomainService) Insert(inputDto model.CreateSystemDeptDto) (model.SystemDept, error) {
	var entity model.SystemDept
	entity.SystemDeptBase = inputDto.SystemDeptBase
	entity.ID = utility.UuidCreate()
	if p, err := orm.ISession().Insert(&entity); err != nil {
		return entity, err
	} else if p <= 0 {
		return entity, errors.New("新增部门失败")
	}
	return entity, nil
}

func (s *systemDeptDomainService) Update(inputDto model.UpdateSystemDeptDto) (model.SystemDept, error) {
	var entity model.SystemDept
	if has, err := orm.QSession(true).ID(inputDto.ID).Get(&entity); err != nil {
		return entity, err
	} else if !has {
		return entity, errors.New("该部门ID不存在")
	}

	entity.SystemDeptBase = inputDto.SystemDeptBase
	if p, err := orm.USession(true).AllCols().ID(inputDto.ID).Update(&entity); err != nil {
		return entity, err
	} else if p <= 0 {
		return entity, errors.New("修改部门失败")
	}
	return entity, nil
}

func (s *systemDeptDomainService) Delete(id string) error {
	var entity model.SystemDept
	if has, err := orm.QSession(true).ID(id).Get(&entity); err != nil {
		return err
	} else if !has {
		return errors.New("该部门ID不存在")
	}
	if p, err := orm.DSession(true).ID(id).Update(&entity); err != nil {
		return err
	} else if p <= 0 {
		return errors.New("删除部门失败")
	}
	return nil
}
