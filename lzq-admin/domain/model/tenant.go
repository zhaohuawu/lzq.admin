package model

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"lzq-admin/domain/dto"
)

func (tenant *Tenant) TableName() string {
	return TableTenant
}

// Tenant 租户
type Tenant struct {
	BaseModel                   `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
	Code                        string `xorm:"unique varchar(32) notnull comment('租户编码')"`
	Name                        string `xorm:"comment('租户名称')"`
	Email                       string `xorm:"comment('邮箱')"`
	Password                    string `xorm:"comment('密码')"`
	Status                      string `xorm:"varchar(64) notnull comment('状态')"`
	ActiveTime                  string `xorm:"comment('激活时间')"`
}

// CreateTenantDto 新增租户Dto
type CreateTenantDto struct {
	Code     string `json:"code" binding:"required" label:"租户编码"`
	Name     string `json:"name" binding:"required" label:"用户名"`
	Email    string `json:"email" label:"邮箱"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// UpdateTenantDto 修改租户Dto
type UpdateTenantDto struct {
	dto.BaseDto
	CreateTenantDto
	Status string `json:"status" binding:"required" label:"状态"`
}

// DeleteTenantDto 删除租户Dto
type DeleteTenantDto struct {
	dto.BaseDto
}
