package domainservice

import (
	"errors"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/4
 * @Version 1.0.0
 */
type authRoleDomainService struct{}

var AuthRoleDomainService = authRoleDomainService{}

func (s *authRoleDomainService) Insert(modelDto model.CreateAuthRoleDto) (model.AuthRole, error) {
	var entity model.AuthRole
	var exist bool
	var err error

	exist, err = orm.QSession(true).Where("Name=?", modelDto.Name).Get(&model.AuthRole{})
	if err != nil {
		return entity, err
	}
	if exist {
		return entity, errors.New("角色名称：" + modelDto.Name + " 已存在，请更换")
	}
	entity = model.AuthRole{
		AuthRoleBase: model.AuthRoleBase{
			Name:        modelDto.Name,
			Code:        modelDto.Code,
			Description: modelDto.Description,
			RoleStatus:  domainconsts.RoleStatusEnable,
		},
	}
	entity.ID = utility.UuidCreate()
	_, err = orm.ISession().Insert(&entity)
	return entity, err
}
