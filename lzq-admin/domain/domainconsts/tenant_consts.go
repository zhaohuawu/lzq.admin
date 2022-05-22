package domainconsts

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

const (
	TenantStatusDisable = "Disable"
	TenantStatusEnable  = "Enable"
)

var TenantConstFlags = map[string]string{
	TenantStatusDisable: "停用",
	TenantStatusEnable:  "启用",
}

const (
	SystemUserStatusDisable = "Disable"
	SystemUserStatusEnable  = "Enable"
)

var SystemUserConstFlags = map[string]string{
	SystemUserStatusDisable: "停用",
	SystemUserStatusEnable:  "启用",
}

const (
	SystemUserSexSecrecy = "Secrecy"
	SystemUserSexMale    = "Male"
	SystemUserSexFemale  = "Female"
)

var SystemUserSexConstFlags = map[string]string{
	SystemUserSexSecrecy: "保密",
	SystemUserSexMale:    "男",
	SystemUserSexFemale:  "女",
}

const (
	PermissionGroupView      = "View"
	PermissionGroupOperation = "Operation"
	PermissionGroupAdvance   = "Advance"
)

var PermissionGroupConstFlags = map[string]string{
	PermissionGroupView:      "查看",
	PermissionGroupOperation: "操作",
	PermissionGroupAdvance:   "高级",
}

const (
	RoleStatusDisable = "Disable"
	RoleStatusEnable  = "Enable"
)

var RoleStatusConstFlags = map[string]string{
	RoleStatusDisable: "停用",
	RoleStatusEnable:  "启用",
}

const (
	CommonStatusDisable = "Disable"
	CommonStatusEnable  = "Enable"
)

var CommonStatusConstFlags = map[string]string{
	CommonStatusDisable: "停用",
	CommonStatusEnable:  "启用",
}

const (
	SysFileStatusUnused   = "Unused"
	SysFileStatusInuse    = "Inuse"
	SysFileStatusErasable = "Erasable"
	SysFileStatusDeleted  = "Deleted"
)

var SysFileStatusConstFlags = map[string]string{
	SysFileStatusUnused:   "未使用",
	SysFileStatusInuse:    "使用中",
	SysFileStatusErasable: "可删除",
	SysFileStatusDeleted:  "已删除",
}
