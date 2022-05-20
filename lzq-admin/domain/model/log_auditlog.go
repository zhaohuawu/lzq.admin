package model

import (
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/16
 * @Version 1.0.0
 */

func (LogAuditLog *LogAuditLog) TableName() string {
	return TableLogAuditLog
}

type LogAuditLog struct {
	ID                          string    `json:"id" xorm:"pk char(36) comment('主键')"`                 //主键
	ApplicationCode             string    `json:"applicationCode" xorm:"varchar(32) comment('应用编码')"`  //应用编码
	UserID                      string    `json:"userId" xorm:"char(36) comment('用户ID')"`              //用户ID
	UserName                    string    `json:"userName" xorm:"varchar(32) comment('用户名')"`          //用户名
	LoginName                   string    `json:"loginName" xorm:"varchar(32) comment('登录名')"`         //登录名
	TenantName                  string    `json:"tenantName" xorm:"comment('租户名称')"`                   //租户名称
	ExecutionTime               time.Time `json:"executionTime" xorm:"notnull comment('执行时间')"`        //执行时间
	ClientIPAddress             string    `json:"clientIpAddress" xorm:"varchar(20) comment('客户端IP')"` //客户端IP
	BrowserInfo                 string    `json:"browserInfo"  xorm:"varchar(500) comment('浏览器信息')"`   //浏览器信息
	HTTPMethod                  string    `json:"httpMethod"  xorm:"varchar(10) comment('HTTP请求方式')"`  //HTTP请求方式
	URL                         string    `json:"url"  xorm:"varchar(500) comment('接口URL')"`           //接口URL
	Exceptions                  string    `json:"exceptions"  xorm:"text comment('异常详情')"`             //异常详情
	Comments                    string    `json:"comments" xorm:"comment('注释')"`                       //注释
	HTTPStatusCode              int       `json:"httpStatusCode" xorm:"comment('HTTP状态码')"`            //HTTP状态码
	EventCode                   string    `json:"eventTag" xorm:"varchar(32) comment('事件编码')"`         //事件编码
	EventName                   string    `json:"eventTagName" xorm:"varchar(32) comment('事件名称')"`     //事件名称
	TenantBaseModel             `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
	LogEntityChanges            []LogEntityChanges `json:"entityChanges"` //更改的实体
}

func (LogEntityChanges *LogEntityChanges) TableName() string {
	return TableLogEntityChanges
}

type LogEntityChanges struct {
	ID                          string    `json:"id" xorm:"pk char(36) comment('主键')"`                   //主键
	AuditLogID                  string    `json:"auditLogId"  xorm:"char(36) comment('审计日志ID')"`         //审计日志ID
	ChangeTime                  time.Time `json:"changeTime" xorm:"notnull comment('更改时间')"`             //更改时间
	ChangeType                  string    `json:"changeType" xorm:"varchar(20) notnull comment('更改类型')"` //更改类型
	EntityId                    string    `json:"entityId" xorm:"char(36) notnull comment('更改实体ID')"`    //更改实体ID
	EntityTypeFullName          string    `json:"entityTypeFullName" xorm:"notnull comment('更改实体全名')"`   //更改实体全名
	TenantBaseModel             `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
	LogEntityPropertyChanges    []LogEntityPropertyChanges `json:"entityPropertyChanges"` //实体更改的属性
}

func (LogEntityPropertyChanges *LogEntityPropertyChanges) TableName() string {
	return TableLogEntityPropertyChanges
}

type LogEntityPropertyChanges struct {
	ID                          string `json:"id" xorm:"pk char(36) comment('主键')"`                          //主键
	EntityChangeId              string `json:"entityChangeId" xorm:"char(36) comment('实体更改ID')"`             //实体更改ID
	NewValue                    string `json:"newValue" xorm:"varchar(500) comment('更改的新值')"`                //更改的新值
	OriginalValue               string `json:"originalValue" xorm:"varchar(500) comment('更改的原值')"`           //更改的原值
	PropertyName                string `json:"propertyName" xorm:"varchar(32) comment('更改的属性')"`             //更改的属性
	PropertyNameDesc            string `json:"propertyNameDesc" xorm:"varchar(100) comment('更改的属性描述')"`      //更改的属性描述
	PropertyTypeFullName        string `json:"propertyTypeFullName" xorm:"varchar(32) comment('更改的属性数据类型')"` //更改的属性数据类型
	TenantBaseModel             `xorm:"extends"`
	HasExtraPropertiesBaseModel `xorm:"extends"`
}

type CreateLogAuditLogDto struct {
	LogAuditLog
}

type LogAuditLogDto struct {
	LogAuditLog `xorm:"extends"`
}
