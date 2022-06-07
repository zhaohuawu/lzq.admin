package model

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/23
 * @Version 1.0.0
 */

func (SystemCompany *SystemCompany) TableName() string {
	return TableSystemCompany
}

type SystemCompany struct {
	BaseModel         `xorm:"extends"`
	TenantBaseModel   `xorm:"extends"`
	SystemCompanyBase `xorm:"extends"`
}

type SystemCompanyBase struct {
	Name     string `json:"name" binding:"required" xorm:"varchar(32) notnull comment('公司名称')"` //公司名称
	ParentId string `json:"parentId" xorm:"char(36) comment('上级公司ID')"`                         //上级公司ID
	Remark   string `json:"remark" xorm:"varchar(200) comment('备注')"`                           //备注
	Rank     int    `json:"rank" xorm:"comment('排序')"`                                          //排序
}

type CreateSystemCompanyDto struct {
	SystemCompanyBase
}

type UpdateSystemCompanyDto struct {
	ID string `json:"id" binding:"required"` //公司ID
	SystemCompanyBase
}

type SystemCompanyDto struct {
	SystemCompany `xorm:"extends"`
}

type SystemCompanyListDto struct {
	SystemCompany `xorm:"extends"`
	Operation     string `json:"operation"` //操作
}
