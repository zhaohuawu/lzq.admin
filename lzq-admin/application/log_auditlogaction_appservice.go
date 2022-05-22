package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/model"
	token "lzq-admin/pkg/auth"
	"lzq-admin/pkg/orm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/19
 * @Version 1.0.0
 */

type logAuditLogActionAppService struct {
	BaseAppService
}

var LogAuditLogActionAppService = logAuditLogActionAppService{}

// GetList doc
// @Summary 查询接口日志列表
// @Tags AuditLogAction
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.LogAuditLogActionListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/auditLogAction/list [GET]
func (app *logAuditLogActionAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result = make([]model.LogAuditLogActionListDto, 0)
	dbSession := orm.QSession(false).Omit("Operation")
	if err := DBCondition(inputDto, dbSession, "", model.LogAuditLogActionListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var resultDto PageListDto
	var err error
	if inputDto.RequireTotalCount {
		resultDto.TotalCount, err = dbSession.FindAndCount(&result)
	} else {
		err = dbSession.Find(&result)
	}
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	resultDto.Data = result
	app.ResponseSuccess(c, resultDto)
}

// GetCurrentUserLogsList doc
// @Summary 查询当前用户的登录日志列表
// @Tags AuditLogAction
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.LogAuditLogActionListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/auditLogAction/currentUserLogsList [GET]
func (app *logAuditLogActionAppService) GetCurrentUserLogsList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	inputDto.Filter = "[[\"userId\",\"=\",\"" + token.GlobalTokenClaims.Id + "\"],[\"actionType\",\"=\",\"Login\"]]"

	var result = make([]model.LogAuditLogActionListDto, 0)
	dbSession := orm.QSession(false).Omit("Operation")
	if err := DBCondition(inputDto, dbSession, "", model.LogAuditLogActionListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var resultDto PageListDto
	var err error
	if inputDto.RequireTotalCount {
		resultDto.TotalCount, err = dbSession.FindAndCount(&result)
	} else {
		err = dbSession.Find(&result)
	}
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	resultDto.Data = result
	app.ResponseSuccess(c, resultDto)
}
