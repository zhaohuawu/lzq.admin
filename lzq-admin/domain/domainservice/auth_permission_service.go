package domainservice

import (
	"github.com/pkg/errors"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"xorm.io/xorm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/2
 * @Version 1.0.0
 */

// AuthPermissionDomainService 操作权限领域服务
type authPermissionDomainService struct{}

var AuthPermissionDomainService = authPermissionDomainService{}

func (s *authPermissionDomainService) Insert(dbSession *xorm.Session, modelDto model.CreateAuthPermissionDto) (result model.AuthPermission, err error) {
	var entity model.AuthPermission
	entity = model.AuthPermission{
		AuthPermissionBase: model.AuthPermissionBase{
			Code:   modelDto.Code,
			Name:   modelDto.Name,
			Rank:   modelDto.Rank,
			MenuId: modelDto.MenuId,
			Policy: modelDto.Policy,
		},
	}
	entity.ID = utility.UuidCreate()

	_, err = orm.ISessionWithTrans(dbSession).Insert(&entity)
	if err != nil {
		dbSession.Rollback()
		return model.AuthPermission{}, err
	}
	return entity, nil
}

func (s *authPermissionDomainService) Delete(dbSession *xorm.Session, menuId string) error {
	m := model.AuthPermission{}
	if _, err := orm.DSessionWithTrans(false, dbSession).Where("menuId = ?", menuId).Update(&m); err != nil {
		dbSession.Rollback()
		return err
	}
	return nil
}

func (s *authPermissionDomainService) Get(id string) (model.AuthPermissionDto, error) {
	var dto model.AuthPermissionDto
	has, err := orm.QSession(false).ID(id).Where("IsDeleted = ?", false).Get(&dto)
	if err != nil {
		return dto, err
	}
	if has {
		return dto, nil
	} else {
		return dto, errors.New("操作权限不存在或已删除")
	}
}

