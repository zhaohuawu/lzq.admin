package model

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	token "lzq-admin/pkg/auth"
	"reflect"
	"time"
)

const (
	TableSystemConfig             = "sys_config"
	TableSystemUser               = "sys_user"
	TableTenant                   = "sys_tenant"
	TableSystemCompany            = "sys_company"
	TableSystemDept               = "sys_dept"
	TableSystemDictionary         = "sys_dictionary"
	TableAuthModule               = "au_module"
	TableAuthMenu                 = "au_menu"
	TableAuthPermission           = "au_permission"
	TableAuthRole                 = "au_role"
	TableAuthRolePermission       = "au_rolepermission"
	TableAuthUserdataprivilege    = "au_userdataprivilege"
	TableSystemFile               = "sys_file"
	TableLogAuditLog              = "log_auditlog"
	TableLogEntityChanges         = "log_entitychanges"
	TableLogEntityPropertyChanges = "log_entitypropertychanges"
	TableLogAuditLogAction        = "log_auditlogaction"
)

// BaseModel 基类
type BaseModel struct {
	ID                   string    `json:"id" xorm:"pk char(36) comment('主键')"` //主键
	CreationTime         time.Time `json:"-" xorm:"created notnull comment('创建时间')"`
	CreatorId            string    `json:"-" xorm:"char(36) comment('创建人')"`
	LastModificationTime time.Time `json:"-" xorm:"comment('最后修改时间')"`
	LastModifierId       string    `json:"-" xorm:"char(36) comment('最后修改人')"`
	IsDeleted            bool      `json:"-" xorm:"bool default(0) comment('是否已删除')"`
	DeleterId            string    `json:"-" xorm:"char(36) comment('删除人')"`
	DeletionTime         time.Time `json:"-" xorm:"comment('删除时间')"`
}

// TenantBaseModel 租户字段基类
type TenantBaseModel struct {
	TenantId string `json:"-" xorm:"char(36) comment('租户ID')"` //租户ID
}

// HasExtraPropertiesBaseModel 扩展字段基类
type HasExtraPropertiesBaseModel struct {
	ExtraProperties map[string]interface{} `json:"-" xorm:"json longtext comment('扩展字段')"`
}

func BeforeInsert(useMultiTenancy bool, obj interface{}) interface{} {
	immutable := reflect.ValueOf(obj).Elem()
	if (token.GlobalTokenClaims != nil && len(token.GlobalTokenClaims.Id) > 0 && immutable.FieldByName("CreatorId") != reflect.Value{}) {
		immutable.FieldByName("CreatorId").SetString(token.GlobalTokenClaims.Id)
	}
	if (useMultiTenancy && token.GlobalTokenClaims != nil && len(token.GlobalTokenClaims.TenantId) > 0 && immutable.FieldByName("TenantId") != reflect.Value{}) {
		immutable.FieldByName("TenantId").SetString(token.GlobalTokenClaims.TenantId)
	}
	return obj
}

func BeforeUpdate(obj interface{}) (interface{}, bool, bool) {
	immutable := reflect.ValueOf(obj).Elem()
	isModityId := false
	isModityTime := false
	if (immutable.FieldByName("LastModificationTime") != reflect.Value{}) {
		isModityTime = true
		immutable.FieldByName("LastModificationTime").Set(reflect.ValueOf(time.Now()))
	}
	if (immutable.FieldByName("LastModifierId") != reflect.Value{}) {
		isModityId = true
		immutable.FieldByName("LastModifierId").SetString(token.GlobalTokenClaims.Id)
	}
	return obj, isModityId, isModityTime
}

func BeforeDelete(obj interface{}) (interface{}, bool, bool) {
	immutable := reflect.ValueOf(obj).Elem()
	isDeleterId := false
	isDeletionTime := false
	if (immutable.FieldByName("DeletionTime") != reflect.Value{}) {
		immutable.FieldByName("DeletionTime").Set(reflect.ValueOf(time.Now()))
		isDeletionTime = true
	}
	if (immutable.FieldByName("IsDeleted") != reflect.Value{}) {
		immutable.FieldByName("IsDeleted").SetBool(true)
	}
	if (immutable.FieldByName("DeleterId") != reflect.Value{}) {
		immutable.FieldByName("DeleterId").SetString(token.GlobalTokenClaims.Id)
		isDeleterId = true
	}
	return obj, isDeleterId, isDeletionTime
}
