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

type systemCompanyDomainService struct {
}

var SystemCompanyDomainService = systemCompanyDomainService{}

func (s *systemCompanyDomainService) Insert(inputDto model.CreateSystemCompanyDto) (model.SystemCompany, error) {
	var entity model.SystemCompany
	entity.SystemCompanyBase = inputDto.SystemCompanyBase
	entity.ID = utility.UuidCreate()
	if p, err := orm.ISession().Insert(&entity); err != nil {
		return entity, err
	} else if p <= 0 {
		return entity, errors.New("新增公司失败")
	}
	return entity, nil
}

func (s *systemCompanyDomainService) Update(inputDto model.UpdateSystemCompanyDto) (model.SystemCompany, error) {
	var entity model.SystemCompany
	if has, err := orm.QSession(true).ID(inputDto.ID).Get(&entity); err != nil {
		return entity, err
	} else if !has {
		return entity, errors.New("该公司ID不存在")
	}

	entity.SystemCompanyBase = inputDto.SystemCompanyBase
	if p, err := orm.USession(true).AllCols().ID(inputDto.ID).Update(&entity); err != nil {
		return entity, err
	} else if p <= 0 {
		return entity, errors.New("修改公司失败")
	}
	return entity, nil
}

func (s *systemCompanyDomainService) Delete(id string) error {
	var entity model.SystemCompany
	if has, err := orm.QSession(true).ID(id).Get(&entity); err != nil {
		return err
	} else if !has {
		return errors.New("该公司ID不存在")
	}
	if p, err := orm.DSession(true).ID(id).Update(&entity); err != nil {
		return err
	} else if p <= 0 {
		return errors.New("删除公司失败")
	}
	return nil
}
