package domainservice

import (
	"github.com/ahmetb/go-linq/v3"
	"lzq-admin/domain/model"
	"lzq-admin/middleware"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"strings"
	"sync"
	"xorm.io/xorm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/6
 * @Version 1.0.0
 */

type authRolePermissionDomainService struct {
	wg sync.WaitGroup
}

var AuthRolePermissionDomainService = authRolePermissionDomainService{}

func (s *authRolePermissionDomainService) Insert(modelDtoes []model.CreateAuthRolePermissionDto) ([]model.AuthRolePermission, error) {
	var iEntities []model.AuthRolePermission
	var uEntities []model.AuthRolePermission
	var err error
	var rolePermissionList = make([]model.AuthRolePermission, 0)
	roleIds := make([]string, 0)
	linq.From(modelDtoes).SelectT(func(s model.CreateAuthRolePermissionDto) string { return s.RoleId }).Distinct().ToSlice(&roleIds)
	if err = orm.QSession(true).Where("RoleId in(?)", strings.Join(roleIds, ",")).Find(&rolePermissionList); err != nil {
		return nil, err
	}
	for _, v := range modelDtoes {
		m := model.AuthRolePermission{
			AuthRolePermissionBase: model.AuthRolePermissionBase{
				RoleId:       v.RoleId,
				PermissionId: v.PermissionId,
				IsGranted:    v.IsGranted,
			},
		}
		firsts := make([]model.AuthRolePermission, 0)
		linq.From(rolePermissionList).WhereT(func(w model.AuthRolePermission) bool {
			return w.RoleId == v.RoleId && w.PermissionId == v.PermissionId
		}).ToSlice(&firsts)
		if len(firsts) > 0 {
			first := firsts[0]
			first.IsGranted = m.IsGranted
			uEntities = append(uEntities, first)
		} else {
			m.ID = utility.UuidCreate()
			m.TenantId = middleware.TokenClaims.TenantId
			m.CreatorId = middleware.TokenClaims.Id
			iEntities = append(iEntities, m)
		}
	}
	var dbSession *xorm.Session
	// 开启事务
	if dbSession, err = orm.BeginTrans(); err != nil {
		dbSession.Rollback()
		return nil, err
	}
	go func() {
		s.wg.Add(1)
		if len(iEntities) > 0 {
			if _, err := dbSession.Insert(&iEntities); err != nil {
				dbSession.Rollback()
				s.wg.Done()
				return
			}
		}
		s.wg.Done()
	}()

	if len(uEntities) > 0 {
		for _, u := range uEntities {
			if _, err := orm.USessionWithTrans(true, dbSession).ID(u.ID).UseBool("IsGranted").Update(&u); err != nil {
				dbSession.Rollback()
				return nil, err
			}
		}
	}
	s.wg.Wait()
	iEntities = append(iEntities, uEntities...)

	if err := dbSession.Commit(); err != nil {
		dbSession.Rollback()
		return nil, err
	}
	return iEntities, err
}
