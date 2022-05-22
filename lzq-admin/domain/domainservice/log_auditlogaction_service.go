package domainservice

import (
	"lzq-admin/config"
	"lzq-admin/domain/model"
	token "lzq-admin/pkg/auth"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/tenant"
	"lzq-admin/pkg/utility"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/17
 * @Version 1.0.0
 */

type logAuditLogActionService struct {
}

var ILogAuditLogActionService = logAuditLogActionService{}

func (s *logAuditLogActionService) Insert(inputDto model.CreateLogAuditLogActionDto) error {
	var result model.LogAuditLogAction
	result.ID = utility.UuidCreate()
	if token.GlobalTokenClaims != nil && len(token.GlobalTokenClaims.Id) > 0 {
		result.UserID = token.GlobalTokenClaims.Id
		result.UserName = token.GlobalTokenClaims.Name
		result.LoginName = token.GlobalTokenClaims.LoginName
		if config.Config.UseMultiTenancy {
			tenantInfo, _ := tenant.GetTenantById(token.GlobalTokenClaims.TenantId)
			result.TenantName = tenantInfo.Name
		}
	}
	result.ServiceModuleCode = config.Config.ServerConfig.ServiceModuleCode
	result.LogAuditLogActionBase = inputDto.LogAuditLogActionBase
	if _, err := orm.ISession().Insert(&result); err != nil {
		return err
	}
	return nil
}

func (s *logAuditLogActionService) AnonymousInsert(inputDto model.LogAuditLogAction) error {
	inputDto.ID = utility.UuidCreate()
	inputDto.ServiceModuleCode = config.Config.ServerConfig.ServiceModuleCode
	if _, err := orm.ISession().Insert(&inputDto); err != nil {
		return err
	}
	return nil
}

