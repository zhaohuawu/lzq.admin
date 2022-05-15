package model

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/15
 * @Version 1.0.0
 */

func (SystemFile *SystemFile) TableName() string {
	return TableSystemFile
}

type SystemFile struct {
	BaseModel       `xorm:"extends"`
	TenantBaseModel `xorm:"extends"`
	SystemFileBase  `xorm:"extends"`
	Status          string `json:"status" xorm:"varchar(32) notnull comment('状态')"` //状态
}

type SystemFileBase struct {
	ContentType  string `json:"contentType" binding:"required" xorm:"varchar(20) notnull comment('文件类型')"` //文件类型
	OriginalName string `json:"originalName" xorm:"comment('文件原名称')"`                                      //文件原名称
	NewName      string `json:"newName" xorm:"comment('文件新名称')"`                                           //文件新名称
	Extension    string `json:"extension" xorm:"varchar(20) comment('文件新名称')"`                             //文件扩展名
	Size         int64  `json:"size" xorm:"notnull comment('文件大小（单位：B）')"`                                 //文件大小（单位：B）
	Url          string `json:"url" binding:"required" xorm:"varchar(500) notnull comment('文件路径')"`        //文件路径
	ThirdPartyId string `json:"thirdPartyId" xorm:"comment('第三方文件存储ID')"`                                  //第三方文件存储ID
}

type CreateSystemFileDto struct {
	SystemFileBase
}

type SystemFileDto struct {
	ID             string `json:"id"` //用户id
	SystemFileBase `xorm:"extends"`
	Status         string `json:"status"` //状态
}
