package cache

import "fmt"

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/26
 * @Version 1.0.0
 */

type cacheKeyHelper struct{}

var LzqCacheKeyHelper = &cacheKeyHelper{}

func (c *cacheKeyHelper) GetUserGrantedMenusCacheKey(userId string) string {
	return fmt.Sprintf("%v:UserGrantedMenus:%v", LzqCacheHelper.GetCacheVersion(LzqCacheTypeFunctionPrivilege), userId)
}

func (c *cacheKeyHelper) GetRoleGrantedPermissionsCacheKey(roleId string) string {
	return fmt.Sprintf("%v:RoleGrantedPermissions:%v", LzqCacheHelper.GetCacheVersion(LzqCacheTypeRolePrivilege), roleId)
}

func (c *cacheKeyHelper) GetGrantedDataPrivilegeByUserCacheKey(userId string) string {
	return fmt.Sprintf("%v:GrantedDataPrivilegeByUser:%v", LzqCacheHelper.GetCacheVersion(LzqCacheTypeDataPrivilege), userId)
}

func (c *cacheKeyHelper) GetUserGrantedPolicyCacheKey(userId string) string {
	return fmt.Sprintf("%v:UserGrantedPolicy:%v", LzqCacheHelper.GetCacheVersion(LzqCacheTypeFunctionPrivilege), userId)
}

func (c *cacheKeyHelper) GetUserGrantedPermissionsCacheKey(userId string) string {
	return fmt.Sprintf("%v:%v:UserGrantedPermissions:%v", LzqCacheHelper.GetCacheVersion(LzqCacheTypeFunctionPrivilege), LzqCacheHelper.GetCacheVersion(LzqCacheTypeDataPrivilege), userId)
}
