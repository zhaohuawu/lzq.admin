package domainservice

import "lzq-admin/domain/model"

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/16
 * @Version 1.0.0
 */

type logAuditLogService struct {
}

var ILogAuditLogService = logAuditLogService{}

func (s *logAuditLogService) Insert(inputDto model.CreateLogAuditLogDto) (model.LogAuditLogDto, error) {
	var result model.LogAuditLogDto
	return result, nil
}
