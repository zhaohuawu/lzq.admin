package model

/**
 * @author  糊涂的老知青
 * @date    2021/11/25
 * @version 1.0
 */

func (AuthModule *AuthModule) TableName() string {
	return TableAuthModule
}

type AuthModule struct {
	BaseModel      `xorm:"extends"`
	AuthModuleBase `xorm:"extends"`
}

type AuthModuleBase struct {
	Name string `json:"name" binding:"required" xorm:"varchar(50) notnull comment('模块名称')"` //模块名称
	Code string `json:"code" binding:"required" xorm:"varchar(20) notnull comment('模块编码')"` //模块编码
	Rank int    `json:"rank" xorm:"default(1) notnull comment('排序')"`                       //排序
}

type CreateAuthModuleDto struct {
	AuthModuleBase
}

type UpdateAuthModuleDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthModuleBase
}

type AuthModuleDto struct {
	AuthModule `xorm:"extends"`
}

type AuthModuleListDto struct {
	AuthModule `xorm:"extends"`
	Operation  string `json:"operation"` //操作
}
