package model

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/12
 * @Version 1.0.0
 */

func (AuthUserDataPrivilege *AuthUserDataPrivilege) TableName() string {
	return TableAuthUserdataprivilege
}

type AuthUserDataPrivilege struct {
	BaseModel                 `xorm:"extends"`
	TenantBaseModel           `xorm:"extends"`
	AuthUserDataPrivilegeBase `xorm:"extends"`
}

type AuthUserDataPrivilegeBase struct {
	RoleId string `json:"roleId" binding:"required" xorm:"char(36) notnull comment('角色ID')"` //角色ID
	UserId string `json:"userId" binding:"required" xorm:"char(36) notnull comment('用户ID')"` //用户ID
}

type CreateAuthUserDataPrivilegeDto struct {
	AuthUserDataPrivilegeBase
}

type UpdateAuthUserDataPrivilegeDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthUserDataPrivilegeBase
}

type AuthUserDataPrivilegeDto struct {
	AuthUserDataPrivilege `xorm:"extends"`
}
