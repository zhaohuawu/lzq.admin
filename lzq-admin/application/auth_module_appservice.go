package application

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/1
 * @Version 1.0.0
 */

type AuthModuleAppService struct {
	BaseAppService
}

var IAuthModuleAppService = AuthModuleAppService{}

// Create doc
// @Summary 创建服务模块
// @Tags AuthModule
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateAuthModuleDto true " "
// @Success 200 {object} model.AuthModuleDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authModule/create [POST]
func (app *AuthModuleAppService) Create(c *gin.Context) {
	var inputDto model.CreateAuthModuleDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	result, err := domainservice.AuthModuleDomainService.Insert(inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, result)
}

// Get doc
// @Summary 根据模块ID或者模块编码查询模块
// @Tags AuthModule
// @Description id和code必须传一个
// @Accept mpfd
// @Produce  json
// @Param id query string false "模块ID"
// @Param code query string false "模块编码"
// @Success 200 {object} model.AuthModuleDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authModule/get [GET]
func (app *AuthModuleAppService) Get(c *gin.Context) {
	id := c.Query("id")
	code := c.Query("code")
	if result, err := domainservice.AuthModuleDomainService.Get(id, code); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改服务模块
// @Tags AuthModule
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateAuthModuleDto true " "
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authModule/update [PUT]
func (app *AuthModuleAppService) Update(c *gin.Context) {
	var inputDto model.UpdateAuthModuleDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var module model.AuthModule
	var has int64
	var err error
	if has, err = orm.QSession(false).ID(inputDto.Id).Count(&module); err != nil {
		app.ResponseError(c, err)
		return
	}
	if has <= 0 {
		app.ResponseError(c, errors.New("服务模块不存在"))
		return
	}
	module.Name = inputDto.Name
	module.Code = inputDto.Code
	module.Rank = inputDto.Rank

	var updateNum int64
	if updateNum, err = orm.USession(false).AllCols().ID(inputDto.Id).Update(&module); err != nil {
		app.ResponseError(c, err)
		return
	}
	if updateNum < 1 {
		app.ResponseError(c, errors.New("修改失败"))
		return
	}
	app.ResponseSuccess(c, true)
}

// Delete
// @Summary 删除服务模块
// @Tags AuthModule
// @Description
// @Produce  json
// @Param id query string true "模块ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authModule/delete [DELETE]
func (app *AuthModuleAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	if _, err := orm.DSession(false).ID(id).Update(&model.AuthModule{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, true)
}

// GetList doc
// @Summary 查询模块列表
// @Tags AuthModule
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.SystemUserListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authModule/list [GET]
func (app *AuthModuleAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	dbSession := orm.QSession(false)
	if err := DBCondition(inputDto, dbSession, "", model.AuthModuleListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result = make([]model.AuthModuleListDto, 0)
	var resultDto PageListDto
	var err error
	dbSession.Omit("Operation")
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
