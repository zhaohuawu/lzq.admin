package model

/**
 * @Author  糊涂的老知青
 * @Date    2022/7/5
 * @Version 1.0.0
 */

func (SystemDictionary *SystemDictionary) TableName() string {
	return TableSystemDictionary
}

type SystemDictionary struct {
	BaseModel            `xorm:"extends"`
	SystemDictionaryBase `xorm:"extends"`
	Status               string `json:"status" xorm:"varchar(200) comment('状态')"` //状态
}

type SystemDictionaryBase struct {
	DictCode  string `json:"dictCode" binding:"required" xorm:"varchar(200) notnull comment('公司名称')"`  //字典编码
	DictKey   string `json:"dictKey" binding:"required" xorm:"varchar(200) notnull comment('公司名称')"`   //字典Key -1：字典，!=-1:字典数据
	DictValue string `json:"dictValue" binding:"required" xorm:"varchar(200) notnull comment('公司名称')"` //字典Value
	Remark    string `json:"remark" xorm:"varchar(200) comment('备注')"`                                 //备注
	Sort      int    `json:"sort" xorm:"comment('排序')"`                                                //排序
}

type CreateSystemDictDto struct {
	DictCode  string `json:"dictCode" binding:"required" xorm:"varchar(200) notnull comment('公司名称')"`  //字典编码
	DictValue string `json:"dictValue" binding:"required" xorm:"varchar(200) notnull comment('公司名称')"` //字典Value
	Remark    string `json:"remark" xorm:"varchar(200) comment('备注')"`                                 //备注
}

type CreateSystemDictionaryDto struct {
	SystemDictionaryBase
}

type UpdateSystemDictionaryDto struct {
	ID string `json:"id" binding:"required"` //部门id
	SystemDictionaryBase
}

type SystemDictionaryDto struct {
	SystemDictionary `xorm:"extends"`
}

type SystemDictionaryListDto struct {
	SystemDictionary `xorm:"extends"`
	StatusText       string `json:"statusText"` // 状态
	Operation        string `json:"operation"`  //操作
}

type SimpleSystemDictionaryDto struct {
	SystemDictionaryBase `xorm:"extends"`
	Status               string `json:"status"` //状态
}
