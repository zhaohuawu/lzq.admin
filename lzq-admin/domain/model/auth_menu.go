package model

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

func (AuthMenu *AuthMenu) TableName() string {
	return TableAuthMenu
}

type AuthMenu struct {
	BaseModel    `xorm:"extends"`
	AuthMenuBase `xorm:"extends"`
}

type AuthMenuBase struct {
	ModuleId string `json:"moduleId" binding:"required" xorm:"char(36) notnull comment('模块ID')"` //模块ID
	Name     string `json:"name" binding:"required" xorm:"varchar(50) notnull comment('菜单名称')"`  //菜单名称
	Code     string `json:"code" xorm:"varchar(50) comment('菜单编码')"`                             //菜单编码
	Policy   string `json:"policy" xorm:"varchar(200) comment('菜单策略')"`                          //菜单策略
	Rank     int    `json:"rank" xorm:"default(1) notnull comment('排序')"`                        //排序
	Icon     string `json:"icon" xorm:"comment('图标')"`                                           //图标
	ParentId string `json:"parentId" xorm:"char(36) comment('父菜单ID')"`                           //父菜单ID
	IsBranch bool   `json:"isBranch" xorm:"bool default(0) notnull comment('是否是容器')"`            //是否是容器
	IsHidden bool   `json:"isHidden" xorm:"bool default(0) notnull comment('是否隐藏')"`             //是否隐藏
	Path     string `json:"path" xorm:"varchar(200) comment('内部路径')"`                            //内部路径
	Url      string `json:"url" xorm:"varchar(200) comment('外部路径')"`                             //外部路径
}

type CreateAuthMenuDto struct {
	AuthMenuBase
}

type UpdateAuthMenuDto struct {
	Id string `json:"id" binding:"required"` //id
	AuthMenuBase
}

type AuthMenuDto struct {
	AuthMenu `xorm:"extends"`
}

type AuthMenuListDto struct {
	AuthMenu   `xorm:"extends"`
	ModuleName string            `json:"moduleName" tAlias:"module" tField:"Name"` //模块名称
	ModuleCode string            `json:"moduleCode" tAlias:"module" tField:"Code"` //模块编码
	Operation  string            `json:"operation"`                                //操作
	Children   []AuthMenuListDto `json:"children"`                                 //子菜单
}

type UserGrantedMenuDto struct {
	Id           string `json:"id"` //id
	AuthMenuBase `xorm:"extends"`
	Children     []UserGrantedMenuDto `json:"children"` //子菜单
}
