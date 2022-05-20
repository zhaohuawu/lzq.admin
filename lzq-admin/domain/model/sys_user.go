package model

import "time"

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

func (SystemUser *SystemUser) TableName() string {
	return TableSystemUser
}

type SystemUser struct {
	BaseModel                   `xorm:"extends"`
	TenantBaseModel             `xorm:"extends"`
	SystemUserBase              `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
	Status                      string `json:"status" xorm:"varchar(64) notnull comment('状态')"` //状态
}

type SystemUserBase struct {
	LoginName     string    `json:"loginName" binding:"required" xorm:"unique varchar(32) notnull comment('登录名')"` //登录名
	UserName      string    `json:"userName" binding:"required" xorm:"comment('用户名')"`                             //用户名
	Email         string    `json:"email" xorm:"varchar(200) comment('邮箱')"`                                       //邮箱
	Password      string    `json:"password" binding:"required" xorm:"notnull comment('密码')"`                      //密码
	LastIp        string    `json:"lastIp" xorm:"varchar(200) comment('最后登录ip')"`                                  //最后登录ip
	LastLogDate   time.Time `json:"lastLogDate" xorm:"varchar(200) comment('最后登录时间')"`                             //最后登录时间
	HeadImgURL    string    `json:"headImgUrl" xorm:"varchar(200) comment('头像')"`                                  //头像
	Sex           string    `json:"sex" xorm:"varchar(200) comment('性别')"`                                         //性别
	Mobile        string    `json:"mobile" xorm:"varchar(200) comment('手机号码')"`                                    //手机号码
	IsTenantAdmin bool      `json:"isTenantAdmin" xorm:"bool default(0) comment('是否是租户管理员')"`                      //是否是超级管理员
}

type CreateSystemUserDto struct {
	SystemUserBase
	SurePassword string   `json:"surePassword" binding:"required"` //确认密码
	RoleIds      []string `json:"roleIds"`                         //角色ID
}

type UpdateSystemUserDto struct {
	ID         string   `json:"id" binding:"required"`       //用户id
	UserName   string   `json:"userName" binding:"required"` //用户名
	Email      string   `json:"email"`                       //邮箱
	HeadImgURL string   `json:"headImgUrl"`                  //头像
	Sex        string   `json:"sex"`                         //性别
	Mobile     string   `json:"mobile"`                      //手机号码
	RoleIds    []string `json:"roleIds"`                     //角色ID
}

type SystemUserDto struct {
	SystemUser `xorm:"extends"`
}

type SystemUserInfoDto struct {
	RoleIds       []string `json:"roleIds"`       //角色ID
	RoleName      string   `json:"roleName"`      //角色名称
	Status        string   `json:"status"`        //状态
	LoginName     string   `json:"loginName"`     //登录名
	UserName      string   `json:"userName"`      //用户名称
	HeadImgURL    string   `json:"headImgUrl"`    //头像URL
	HeadImgLink   string   `json:"headImgLink"`   //头像连接
	Sex           string   `json:"sex"`           //性别
	Mobile        string   `json:"mobile"`        //手机号码
	Email         string   `json:"email"`         //邮箱
	IsTenantAdmin bool     `json:"isTenantAdmin"` //是否是租户管理员
	ID            string   `json:"id"`
	SuperAdmin    bool     `json:"superAdmin"` //是否是超级管理员
}

type SystemUserListDto struct {
	SystemUser  `xorm:"extends"`
	StatusName  string   `json:"statusName"`                           //状态
	HeadImgLink string   `json:"headImgLink"`                          //头像连接
	RoleIds     []string `json:"roleIds" tAlias:"urp" tField:"RoleId"` //角色ID
	Operation   string   `json:"operation"`                            //操作
}

type UpdateSystemUserPasswordDto struct {
	ID string `json:"id" binding:"required"`
	UpdateSystemUserPasswordBaseDto
}

type UpdateSystemUserPasswordBaseDto struct {
	Password     string `json:"password" binding:"required"`     //原始密码
	NewPassword  string `json:"newPassword" binding:"required"`  //新密码
	SurePassword string `json:"surePassword" binding:"required"` //确认密码
}
