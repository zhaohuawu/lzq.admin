package domainservice

import (
	"encoding/json"
	"github.com/ahmetb/go-linq/v3"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"strings"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/27
 * @Version 1.0.0
 */

// authCheckerDomainService
type authCheckerDomainService struct {
	wg sync.WaitGroup
}

var AuthCheckerDomainService = authCheckerDomainService{}

func (d *authCheckerDomainService) GetRoleGrantedPermissions(roleId string) ([]dto.RoleGrantPermissionDto, error) {
	result := make([]dto.RoleGrantPermissionDto, 0)
	cacheKey := cache.LzqCacheKeyHelper.GetRoleGrantedPermissionsCacheKey(roleId)
	cacheJson := cache.RedisUtil.NewRedis(true).Get(cacheKey)
	if cacheJson != "" {
		_ = json.Unmarshal([]byte(cacheJson), &result)
		return result, nil
	}
	dbSession := orm.QSession(true, "rp").
		Table(model.TableAuthRolePermission).Alias("rp").
		Join("INNER", model.TableAuthPermission+" as p", orm.ConditionWithDeletedOrTenantId(false, "rp.PermissionId=p.Id", "p")).
		Join("INNER", model.TableAuthMenu+" as m", orm.ConditionWithDeletedOrTenantId(false, "p.MenuId=m.Id", "m")).
		Select("rp.Id,rp.RoleId,CONCAT(m.Policy,':',p.Policy) as Policy,rp.IsGranted").
		Where("rp.RoleId=? and rp.IsGranted=?", roleId, true)
	if err := dbSession.Find(&result); err != nil {
		return result, err
	}

	cache.RedisUtil.NewRedis(true).SSet(cacheKey, result, 0)
	return result, nil
}

func (d *authCheckerDomainService) GetGrantedDataPrivilegesByUser(userId string) ([]dto.UserDataPrivilegeDto, error) {
	result := make([]dto.UserDataPrivilegeDto, 0)
	cacheKey := cache.LzqCacheKeyHelper.GetGrantedDataPrivilegeByUserCacheKey(userId)
	cacheJson := cache.RedisUtil.NewRedis(true).Get(cacheKey)
	if cacheJson != "" {
		_ = json.Unmarshal([]byte(cacheJson), &result)
		return result, nil
	}
	dbSession := orm.QSession(true, "u", "udp", "r").
		Table(model.TableSystemUser).Alias("u").
		Join("INNER", model.TableAuthUserdataprivilege+" as udp", "u.Id=udp.UserId").
		Join("INNER", model.TableAuthRole+" as r", "udp.RoleId=r.Id").
		Select("udp.Id,udp.RoleId,udp.UserId").
		Where("udp.UserId=? and u.Status=? and r.RoleStatus=?", userId, domainconsts.SystemUserStatusEnable, domainconsts.RoleStatusEnable)
	if err := dbSession.Find(&result); err != nil {
		return result, err
	}

	cache.RedisUtil.NewRedis(true).SSet(cacheKey, result, 0)
	return result, nil
}

func (d *authCheckerDomainService) GetUserGrantedRoleIds(userId string) ([]string, error) {
	if userDataPrivileges, err := d.GetGrantedDataPrivilegesByUser(userId); err != nil {
		return nil, err
	} else {
		roleIds := make([]string, 0)
		linq.From(userDataPrivileges).SelectT(func(s dto.UserDataPrivilegeDto) string {
			return s.RoleID
		}).Distinct().ToSlice(&roleIds)
		return roleIds, nil
	}
}

func (d *authCheckerDomainService) IsUserGranted(userId, policy string) bool {
	if len(policy) == 0 {
		return true
	}
	// 租户超级管理员不用校验权限
	if isSuperAdmin, err := SystemUserDomainService.IsSuperAdmin(userId); err != nil {
		return false
	} else if isSuperAdmin {
		return true
	}
	actualPolicy := model.GetActualPolicy(policy)

	cacheKey := cache.LzqCacheKeyHelper.GetUserGrantedPolicyCacheKey(userId)
	cacheJson := cache.RedisUtil.NewRedis(true).HGet(cacheKey, actualPolicy)
	if cacheJson != nil {
		return cacheJson == "1"
	}
	if roleIds, err := d.GetUserGrantedRoleIds(userId); err != nil {
		return false
	} else {
		for _, v := range roleIds {
			roleGrantedPermissions, _ := d.GetRoleGrantedPermissions(v)
			isGranted := linq.From(roleGrantedPermissions).AnyWithT(func(w dto.RoleGrantPermissionDto) bool {
				return strings.ToLower(w.Policy) == strings.ToLower(actualPolicy)
			})
			if isGranted {
				cache.RedisUtil.NewRedis(true).HSet(cacheKey, actualPolicy, true, 0)
				return true
			}
		}
	}
	cache.RedisUtil.NewRedis(true).HSet(cacheKey, actualPolicy, false, 0)
	return false
}

func (d *authCheckerDomainService) GetUserGrantedPermissions(userId string) ([]string, error) {
	result := make([]string, 0)
	cacheKey := cache.LzqCacheKeyHelper.GetUserGrantedPermissionsCacheKey(userId)
	cacheJson := cache.RedisUtil.NewRedis(true).Get(cacheKey)
	if cacheJson != "" {
		_ = json.Unmarshal([]byte(cacheJson), &result)
		return result, nil
	}
	if userGrantedRoleIds, err := d.GetUserGrantedRoleIds(userId); err != nil {
		return nil, err
	} else {
		for _, v := range userGrantedRoleIds {
			if roleGrantedPermissions, err := d.GetRoleGrantedPermissions(v); err != nil {
				return nil, err
			} else if len(roleGrantedPermissions) > 0 {
				p := make([]string, 0)
				linq.From(roleGrantedPermissions).WhereT(func(w dto.RoleGrantPermissionDto) bool {
					return w.IsGranted
				}).SelectT(func(s dto.RoleGrantPermissionDto) string {
					return s.Policy
				}).Distinct().ToSlice(&p)
				if len(p) > 0 {
					result = append(result, p...)
				}
			}
		}
	}
	cache.RedisUtil.NewRedis(true).SSet(cacheKey, result, 0)
	return result, nil
}
