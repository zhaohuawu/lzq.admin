package application

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
)

type TenantAppService struct {
	BaseAppService
}

var ITenantAppService = TenantAppService{}

// Create 测试接口
// @Summary 创建租户
// @Tags Tenant
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateTenantDto true " "
// @Success 200 object model.CreateTenantDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/tenant/create [POST]
func (tenant *TenantAppService) Create(c *gin.Context) {
	var inputDto model.CreateTenantDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		// gin.H封装了生成json数据的工具
		tenant.ResponseError(c, err)
		return
	}

	result, rerr := domainservice.TenantDomainService.Insert(inputDto)
	if rerr != nil {
		// gin.H封装了生成json数据的工具
		tenant.ResponseError(c, rerr)
		return
	}
	fmt.Println(result)
	tenant.ResponseSuccess(c, inputDto)
}
