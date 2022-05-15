package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
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
	result, err := domainservice.SysConfigDomainService.Insert(inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	app.ResponseSuccess(c, result)
}

// GetInfo doc
// @Summary 七牛云配置详情
// @Tags SysConfig
// @Description
// @Accept mpfd
// @Produce  json
// @Param configType query string true "配置类型"
// @Param code query string true "配置编码"
// @Success 200 {object} model.QiNiuConfigDto " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/getInfo [GET]
func (app *systemConfigAppService) GetInfo(c *gin.Context) {
	code := c.Query("code")
	configType := c.Query("configType")
	var result model.SystemConfigDto
	sysConfig, err := domainservice.SysConfigDomainService.GetByCode(configType, code)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	result.ID = sysConfig.ID
	result.ConfigType = sysConfig.ConfigType
	result.Code = sysConfig.Code
	result.Name = sysConfig.Name
	result.ExtraValue = sysConfig.ExtraProperties[model.ExtraSysConfigKey]

	app.ResponseSuccess(c, result)
}

// GetSysConfigJsonMapCache doc
// @Summary 设置系统配置Json及注解缓存
// @Tags SysConfig
// @Description
// @Produce  json
// @Param configType query string true "配置类型"
// @Success 200 {object} interface{} “ ”
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/getSysConfigCache [GET]
func (app *systemConfigAppService) GetSysConfigJsonMapCache(c *gin.Context) {
	configType := c.Query("configType")
	objJson := domainservice.SysConfigDomainService.SetSysConfigJsonMapCache(configType)
	app.ResponseSuccess(c, objJson)
}

// QiuNiuUpdate doc
// @Summary 配置七牛云
// @Tags SysConfig
// @Description
// @Accept json
// @Produce  json
// @Param object body model.QiNiuConfigDto true " "
// @Success 200 {object} model.SystemConfig " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysconfig/qnupdate [POST]
func (app *systemConfigAppService) QiuNiuUpdate(c *gin.Context) {
	var inputDto model.QiNiuConfigDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	sysConfig := model.UpdateSystemConfigDto{
		SystemConfigBase: model.SystemConfigBase{
			//Name:       inputDto.SystemConfigBase.Name,
			ConfigType: inputDto.ConfigUpdateDtoBase.ConfigType,
			Code:       inputDto.ConfigUpdateDtoBase.Code,
		},
		ExtraValue: inputDto.ExtraValue,
	}
	_, err := domainservice.SysConfigDomainService.Update(sysConfig)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	app.ResponseSuccess(c, true)
}
