package model

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/4
 * @Version 1.0.0
 */

func (AuthRolePermission *AuthRolePermission) TableName() string {
	return TableAuthRolePermission
}

type AuthRolePermission struct {
	BaseModel              `xorm:"extends"`
	TenantBaseModel        `xorm:"extends"`
	AuthRolePermissionBase `xorm:"extends"`
}

type AuthRolePermissionBase struct {
	RoleId       string `json:"roleId" binding:"required" xorm:"char(36) notnull comment('角色ID')"`         //角色ID
	PermissionId string `json:"permissionId" binding:"required" xorm:"char(36) notnull comment('操作权限ID')"` //操作权限ID
	IsGranted    bool   `json:"isGranted" xorm:"bool default(1) notnull comment('是否已授权')"`                 //是否已授权
}

type CreateAuthRolePermissionDto struct {
	AuthRolePermissionBase
}

type UpdateAuthRolePermissionDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthRolePermissionBase
}

type AuthRolePermissionDto struct {
	AuthRolePermission `xorm:"extends"`
}

type AuthRolePermissionListDto struct {
	AuthRolePermission `xorm:"extends"`
	Operation          string `json:"operation"` //操作
}
