package domainservice

import "lzq-admin/pkg/cache"

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/28
 * @Version 1.0.0
 */

type authPrivilegeCacheService struct {
}

var AuthPrivilegeCacheService = authPrivilegeCacheService{}

func (s *authPrivilegeCacheService) DeleteFunctionPrivilegeCache() {
	cache.RedisUtil.NewRedis(false).Delete(cache.LzqCacheHelper.GetCacheTypeVersionKey(cache.LzqCacheTypeFunctionPrivilege))
}

func (s *authPrivilegeCacheService) DeleteRoleGrantedPermissionsCache(roleIds ...string) {
	for _, v := range roleIds {
		cache.RedisUtil.NewRedis(true).Delete(cache.LzqCacheKeyHelper.GetRoleGrantedPermissionsCacheKey(v))
	}
}

func (s *authPrivilegeCacheService) DeleteDataPrivilegeCache() {
	cache.RedisUtil.NewRedis(true).Delete(cache.LzqCacheHelper.GetCacheTypeVersionKey(cache.LzqCacheTypeDataPrivilege))
}

func (s *authPrivilegeCacheService) DeleteUserAuthCache(userId string) {
	cache.RedisUtil.NewRedis(true).Delete(cache.LzqCacheKeyHelper.GetUserGrantedMenusCacheKey(userId))
	cache.RedisUtil.NewRedis(true).Delete(cache.LzqCacheKeyHelper.GetGrantedDataPrivilegeByUserCacheKey(userId))
	cache.RedisUtil.NewRedis(true).Delete(cache.LzqCacheKeyHelper.GetUserGrantedPermissionsCacheKey(userId))
}
