package model

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/4
 * @Version 1.0.0
 */

func (AuthRole *AuthRole) TableName() string {
	return TableAuthRole
}

type AuthRole struct {
	BaseModel       `xorm:"extends"`
	TenantBaseModel `xorm:"extends"`
	AuthRoleBase    `xorm:"extends"`
}

type AuthRoleBase struct {
	Name        string `json:"name" binding:"required" xorm:"varchar(100) notnull comment('角色名称')"` //角色名称
	Code        string `json:"code"  xorm:"varchar(200) notnull comment('角色名称')"`                   //角色编码
	Description string `json:"description" xorm:"varchar(200) comment('备注')"`                       //备注
	RoleStatus  string `json:"roleStatus" xorm:"varchar(200) notnull comment('角色状态')"`              //角色状态
}

type CreateAuthRoleDto struct {
	AuthRoleBase
}

type UpdateAuthRoleDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthRoleBase
}

type AuthRoleDto struct {
	AuthRole `xorm:"extends"`
}

type AuthRoleListDto struct {
	AuthRole          `xorm:"extends"`
	RoleStatusText    string          `json:"roleStatusText"`    //角色状态
	HaveAuthorizeUser []AuthorizeUser `json:"haveAuthorizeUser"` //角色授权的用户
	Operation         string          `json:"operation"`         //操作
}

type AuthRoleSimpleDto struct {
	Id          string `json:"id"`          //id
	Name        string `json:"name"`        //角色名称
	Code        string `json:"code"`        //角色编码
	Description string `json:"description"` //备注
}

type AuthorizeUser struct {
	RoleId              string `json:"roleId"`              // 角色ID
	UserId              string `json:"userId"`              //用户ID
	UserName            string `json:"userName"`            // 用户名
	LoginName           string `json:"loginName"`           // 登录名
	UserDataPrivilegeID string `json:"userDataPrivilegeId"` // 用户授权ID
	IsCanDelete         bool   `json:"isCanDelete"`         //是否可以删除
	IsTenantAdmin       bool   `json:"isTenantAdmin"`       // 是否是租户管理员
}
