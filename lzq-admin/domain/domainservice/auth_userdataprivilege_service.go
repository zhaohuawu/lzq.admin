package domainservice

import (
	"github.com/ahmetb/go-linq/v3"
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

func (s *authUserdataPrivilegeDomainService) Insert(dbSession *xorm.Session, userId string, roleIds []string) (result []model.AuthUserDataPrivilege, err error) {
	entities := make([]model.AuthUserDataPrivilege, 0)
	if err := orm.QSession(true).Where("UserId = ?", userId).Find(&entities); err != nil {
		return nil, err
	}
	// 新增
	for _, roleId := range roleIds {
		//var exitEntity model.AuthUserDataPrivilege
		exitIndex := linq.From(entities).IndexOfT(func(w model.AuthUserDataPrivilege) bool {
			return w.RoleId == roleId
		})
		// 将已经存在的过滤掉，没有选择的则是需要删除
		if exitIndex >= 0 {
			entities = append(entities[:exitIndex], entities[exitIndex+1:]...)
			continue
		}
		entity := model.AuthUserDataPrivilege{
			AuthUserDataPrivilegeBase: model.AuthUserDataPrivilegeBase{
				RoleId: roleId,
				UserId: userId,
			},
		}
		entity.ID = utility.UuidCreate()
		_, err = orm.ISessionWithTrans(dbSession).Insert(&entity)
		if err != nil {
			dbSession.Rollback()
			return nil, err
		}
		result = append(result, entity)
	}

	for _, ud := range entities {
		_, err = orm.DSessionWithTrans(true, dbSession).ID(ud.ID).Update(&model.AuthUserDataPrivilege{})
		if err != nil {
			dbSession.Rollback()
			return nil, err
		}
	}
	if len(result) > 0 {
		AuthPrivilegeCacheService.DeleteUserAuthCache(userId)
	}
	return result, nil
}

func (s *authUserdataPrivilegeDomainService) Delete(dbSession *xorm.Session, userId string, roleId string) error {
	m := model.AuthUserDataPrivilege{}
	if len(roleId) > 0 {
		if _, err := orm.DSessionWithTrans(true, dbSession).Where("UserId = ? and RoleId=?", userId, roleId).Update(&m); err != nil {
			dbSession.Rollback()
			return err
		}
	} else {
		if _, err := orm.DSessionWithTrans(true, dbSession).Where("UserId = ?", userId).Update(&m); err != nil {
			dbSession.Rollback()
			return err
		}
	}
	return nil
}
