package model

import "time"

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/17
 * @Version 1.0.0
 */

func (LogAuditLog *LogAuditLogAction) TableName() string {
	return TableLogAuditLogAction
}

type LogAuditLogAction struct {
	ID                    string `json:"id" xorm:"pk char(36) comment('主键')"`                    //主键
	ServiceModuleCode     string `json:"serviceModuleCode" xorm:"varchar(32) comment('服务模块编码')"` //服务模块编码
	UserID                string `json:"userId" xorm:"char(36) comment('用户ID')"`                 //用户ID
	UserName              string `json:"userName" xorm:"varchar(32) comment('用户名')"`             //用户名
	LoginName             string `json:"loginName" xorm:"varchar(32) comment('登录名')"`            //登录名
	TenantName            string `json:"tenantName" xorm:"comment('租户名称')"`                      //租户名称
	IsDeleted             bool   `json:"-" xorm:"bool default(0) comment('是否已删除')"`
	TenantBaseModel       `xorm:"extends"`
	LogAuditLogActionBase `xorm:"extends"`
}

type LogAuditLogActionBase struct {
	ActionType        string                 `json:"actionType"  xorm:"varchar(32) comment('操作类型')"`       //操作类型
	FromSource        string                 `json:"fromSource"  xorm:"varchar(32) comment('来源')"`         //来源
	ClientIPAddress   string                 `json:"clientIpAddress" xorm:"varchar(20) comment('客户端IP')"`  //客户端IP
	BrowserInfo       string                 `json:"browserInfo"  xorm:"varchar(500) comment('浏览器信息')"`    //浏览器信息
	HTTPMethod        string                 `json:"httpMethod"  xorm:"varchar(10) comment('HTTP请求方式')"`   //HTTP请求方式
	URL               string                 `json:"url"  xorm:"varchar(500) comment('接口URL')"`            //接口URL
	Parameters        string                 `json:"parameters"  xorm:"varchar(1000) comment('接口参数')"`     //接口参数
	ExecutionTime     time.Time              `json:"executionTime" xorm:"notnull comment('执行时间')"`         //执行时间
	ExecutionDuration int64                  `json:"executionDuration" xorm:"notnull comment('执行时长')"`     //执行时长
	Exceptions        string                 `json:"exceptions"  xorm:"text comment('异常详情')"`              //异常详情
	Comments          string                 `json:"comments" xorm:"comment('注释')"`                        //注释
	HTTPStatusCode    int                    `json:"httpStatusCode" xorm:"comment('HTTP状态码')"`             //HTTP状态码
	ExtraProperties   map[string]interface{} `json:"extraProperties" xorm:"json longtext comment('扩展字段')"` //扩展字段
}

type CreateLogAuditLogActionDto struct {
	LogAuditLogActionBase
}

type LogAuditLogActionListDto struct {
	LogAuditLogAction `xorm:"extends"`
	Operation         string `json:"operation"` //操作
}

type LogAuditLogActionDto struct {
	LogAuditLogAction `xorm:"extends"`
	ExtraProperty     map[string]interface{} `json:"extraProperty"` //扩展字段
}
