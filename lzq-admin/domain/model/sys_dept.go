package model

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/23
 * @Version 1.0.0
 */

func (SystemDept *SystemDept) TableName() string {
	return TableSystemDept
}

type SystemDept struct {
	BaseModel       `xorm:"extends"`
	TenantBaseModel `xorm:"extends"`
	SystemDeptBase  `xorm:"extends"`
}

type SystemDeptBase struct {
	Name      string `json:"name" binding:"required" xorm:"varchar(32) notnull comment('公司名称')"` //部门名称
	CompanyId string `json:"companyId" xorm:"char(36) comment('所属公司ID')"`                        //所属公司ID
	ParentId  string `json:"parentId" xorm:"char(36) comment('上级部门ID')"`                         //上级部门ID
	Remark    string `json:"remark" xorm:"varchar(200) comment('备注')"`                           //备注
	Rank      int    `json:"rank" xorm:"comment('排序')"`                                          //排序
}

type CreateSystemDeptDto struct {
	SystemDeptBase
}

type UpdateSystemDeptDto struct {
	ID string `json:"id" binding:"required"` //部门id
	SystemDeptBase
}

type SystemDeptDto struct {
	SystemDept `xorm:"extends"`
}

type SystemDeptListDto struct {
	SystemDept `xorm:"extends"`
	Operation  string `json:"operation"` //操作
}

type SystemCompanyAndDeptListDto struct {
	ID        string                        `json:"id"`        //公司或部门id
	Name      string                        `json:"name"`      //公司名称
	CompanyId string                        `json:"companyId"` //所属公司ID
	ParentId  string                        `json:"parentId"`  //上级部门ID
	Remark    string                        `json:"remark"`    //备注
	Rank      int                           `json:"rank"`      //排序
	Type      string                        `json:"type"`      //数据类型，Company：公司，Dept：部门
	Children  []SystemCompanyAndDeptListDto `json:"children"`  //子公司或部门
	Operation string                        `json:"operation"` //操作
}
