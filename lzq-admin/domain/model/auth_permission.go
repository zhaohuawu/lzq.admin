package model

import (
	"fmt"
	"strings"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

func (AuthPermission *AuthPermission) TableName() string {
	return TableAuthPermission
}

type AuthPermission struct {
	BaseModel          `xorm:"extends"`
	AuthPermissionBase `xorm:"extends"`
}

type AuthPermissionBase struct {
	PermissionGroup string `json:"permissionGroup" binding:"required" xorm:"varchar(50) comment('权限组')"`  //权限组
	MenuId          string `json:"menuId" binding:"required" xorm:"char(36) notnull comment('菜单ID')"`     //菜单ID
	Name            string `json:"name" binding:"required" xorm:"varchar(50) notnull comment('权限名称')"`    //权限名称
	Code            string `json:"code" xorm:"varchar(50) comment('权限编码')"`                               //权限编码
	Policy          string `json:"policy" binding:"required" xorm:"varchar(200) notnull comment('权限策略')"` //权限策略
	Rank            int    `json:"rank" xorm:"default(1) notnull comment('排序')"`                          //排序
}

type CreateAuthPermissionDto struct {
	AuthPermissionBase
}

type UpdateAuthPermissionDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthPermissionBase
}

type AuthPermissionDto struct {
	AuthPermission `xorm:"extends"`
}

type AuthPermissionListDto struct {
	AuthPermission      `xorm:"extends"`
	PermissionGroupText string `json:"permissionGroupText"` //权限组
	TypeCode            string `json:"typeCode"`            //类型
	//MenuPolicy     string                  `json:"-"`                                  //模块名称
	MenuName     string                  `json:"menuName"  tAlias:"m" tField:"Name"` //模块名称
	ParentId     string                  `json:"parentId"`                           //父ID
	IsBranch     bool                    `json:"-"`                                  //是否是容器
	ActualPolicy string                  `json:"actualPolicy"`                       //权限策略
	Operation    string                  `json:"operation"`                          //操作
	Children     []AuthPermissionListDto `json:"children"`                           //子菜单
}

func GetActualPolicy(policy string) string {
	if len(policy) == 0 {
		return ""
	}
	actualPolicy := policy
	if !strings.Contains(policy, ":") {
		actualPolicy = fmt.Sprintf("%v:%v", policy, DefaultViewPolicy())
	}
	return actualPolicy
}

func DefaultViewPolicy() string {
	return "View.Access"
}