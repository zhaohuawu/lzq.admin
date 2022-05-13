package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
	"lzq-admin/domain/model/extrastruct"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/4/5
 * @Version 1.0.0
 */

type systemConfigAppService struct {
	BaseAppService
	wg sync.WaitGroup
}

var ISystemConfigAppService = systemConfigAppService{}

// Create doc
// @Summary 系统配置创建
// @Tags SysConfig
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemConfigDto true " "
// @Success 200 {object} model.SystemConfig " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/create [POST]
func (app *systemConfigAppService) Create(c *gin.Context) {
	var inputDto model.CreateSystemConfigDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	_, err := domainservice.SysConfigDomainService.Insert(inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	app.ResponseSuccess(c, true)
}

// QiuNiuUpdate doc
// @Summary 配置七牛云
// @Tags SysConfig
// @Description
// @Accept json
// @Produce  json
// @Param object body extrastruct.QiNiuConfigDto true " "
// @Success 200 {object} model.SystemConfig " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/qnupdate [POST]
func (app *systemConfigAppService) QiuNiuUpdate(c *gin.Context) {
	var inputDto extrastruct.QiNiuConfigDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	sysConfig := model.UpdateSystemConfigDto{
		SystemConfigBase: model.SystemConfigBase{
			Name:       inputDto.SystemConfigBase.Name,
			ConfigType: inputDto.SystemConfigBase.ConfigType,
			Code:       inputDto.SystemConfigBase.Code,
		},
		ExtraValue: inputDto.ExtraQiNiuConfig,
	}
	_, err := domainservice.SysConfigDomainService.Update(sysConfig)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	app.ResponseSuccess(c, true)
}

// GetQiuNiuInfo doc
// @Summary 七牛云配置详情
// @Tags SysConfig
// @Description
// @Accept mpfd
// @Produce  json
// @Param id query string true "配置ID"
// @Success 200 {object} extrastruct.QiNiuConfigDto " "
// @Failure 500 {object} extrastruct.QiNiuConfigDto
// @Router /api/app/sysconfig/getQnInfo [GET]
func (app *systemConfigAppService) GetQiuNiuInfo(c *gin.Context) {
	id := c.Query("id")

	var result extrastruct.QiNiuConfigDto
	if len(id) > 0 {
		sysConfig, err := domainservice.SysConfigDomainService.GetById(id)
		if err != nil {
			app.ResponseError(c, err)
			return
		}
		result.Id = sysConfig.ID
		result.ConfigType = sysConfig.ConfigType
		result.Code = sysConfig.Code
		result.ExtraQiNiuConfig = sysConfig.ExtraProperties[model.ExtraSysConfigKey].(extrastruct.ExtraQiNiuConfig)
	}
	app.ResponseSuccess(c, result)
}

// GetSysconfigJsonMapCache doc
// @Summary 设置系统配置Json及注解缓存
// @Tags SysConfig
// @Description
// @Produce  json
// @Param configType query string true "配置类型"
// @Success 200 {object} map “ ”
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/getSysConfigCache [GET]
func (app *systemConfigAppService) GetSysconfigJsonMapCache(c *gin.Context) {
	configType := c.Query("configType")
	objJson:= domainservice.SysConfigDomainService.SetSysconfigJsonMapCache(configType)
	app.ResponseSuccess(c, objJson)
}
