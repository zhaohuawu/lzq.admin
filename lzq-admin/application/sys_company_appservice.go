package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/23
 * @Version 1.0.0
 */

type systemCompanyAppService struct {
	BaseAppService
}

var SystemCompanyAppService = systemCompanyAppService{}

// Create doc
// @Summary 新增公司
// @Tags SystemCompany
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemCompanyDto true " "
// @Success 200 {object} model.SystemCompany 其他描述
// @Failure 500 {object} ResponseDto
// @Router /api/app/company/create [POST]
func (app *systemCompanyAppService) Create(c *gin.Context) {
	var inputDto model.CreateSystemCompanyDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemCompanyDomainService.Insert(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改公司
// @Tags SystemCompany
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemCompanyDto true " "
// @Success 200 {object} model.SystemCompany 其他描述
// @Failure 500 {object} ResponseDto
// @Router /api/app/company/update [PUT]
func (app *systemCompanyAppService) Update(c *gin.Context) {
	var inputDto model.UpdateSystemCompanyDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemCompanyDomainService.Update(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Delete
// @Summary 删除公司
// @Tags SystemCompany
// @Description
// @Produce  json
// @Param id query string true "公司ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/company/delete [DELETE]
func (app *systemCompanyAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	if err := domainservice.SystemCompanyDomainService.Delete(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}
