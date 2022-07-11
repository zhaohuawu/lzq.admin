package cache

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/25
 * @Version 1.0.0
 */

const (
	LzqCacheTypeSysUser           = "CacheTypeSysUser"
	LzqCacheTypeFunctionPrivilege = "CacheTypeFunctionPrivilege"
	LzqCacheTypeRolePrivilege     = "CacheTypeRolePrivilege"
	LzqCacheTypeDataPrivilege     = "CacheTypeDataPrivilege"
	LzqCacheTypeSystemDict        = "CacheTypeSystemDict"
)

var LzqCacheTypeConstFlags = map[string]string{
	LzqCacheTypeSysUser:           "系统用户缓存类型",
	LzqCacheTypeFunctionPrivilege: "功能权限缓存类型",
	LzqCacheTypeRolePrivilege:     "角色权限缓存类型",
	LzqCacheTypeDataPrivilege:     "数据权限缓存类型",
	LzqCacheTypeSystemDict:        "系统字典缓存类型",
}
