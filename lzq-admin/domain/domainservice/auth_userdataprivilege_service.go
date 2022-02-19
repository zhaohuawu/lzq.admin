package domainservice

import (
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"xorm.io/xorm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/16
 * @Version 1.0.0
 */

// AuthPermissionDomainService 操作权限领域服务
type authUserdataPrivilegeDomainService struct{}

var AuthUserdataPrivilegeDomainService = authUserdataPrivilegeDomainService{}

func (s *authUserdataPrivilegeDomainService) Insert(dbSession *xorm.Session, modelDto model.CreateAuthUserDataPrivilegeDto) (result model.AuthUserDataPrivilege, err error) {
	var entity model.AuthUserDataPrivilege
	if has, err := orm.QSession(true).Where("RoleId = ? and UserId = ?", modelDto.RoleId, modelDto.UserId).Get(&entity); err != nil {
		return entity, err
	} else if has {
		return entity, nil
	}

	entity = model.AuthUserDataPrivilege{
		AuthUserDataPrivilegeBase: model.AuthUserDataPrivilegeBase{
			RoleId: modelDto.RoleId,
			UserId: modelDto.UserId,
		},
	}

	if _, err := orm.DSessionWithTrans(true, dbSession).Where("UserId = ?", modelDto.UserId).Update(&model.AuthUserDataPrivilege{}); err != nil {
		dbSession.Rollback()
		return model.AuthUserDataPrivilege{}, err
	}

	entity.ID = utility.UuidCreate()
	_, err = orm.ISessionWithTrans(dbSession).Insert(&entity)
	if err != nil {
		dbSession.Rollback()
		return model.AuthUserDataPrivilege{}, err
	}

	return entity, nil
}

func (s *authUserdataPrivilegeDomainService) Delete(dbSession *xorm.Session, userId string, roleId string) error {
	m := model.AuthUserDataPrivilege{}
	if _, err := orm.DSessionWithTrans(true, dbSession).Where("UserId = ?", userId).Update(&m); err != nil {
		dbSession.Rollback()
		return err
	}
	return nil
}
