package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/7/6
 * @Version 1.0.0
 */

type systemDictionaryAppService struct {
	BaseAppService
}

var SystemDictionaryAppService = systemDictionaryAppService{}

// CreateDict doc
// @Summary 创建字典父级
// @Tags SystemDictionary
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemDictDto true " "
// @Success 200 {object} model.SystemDictionary " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/createDict [POST]
func (app *systemDictionaryAppService) CreateDict(c *gin.Context) {
	var inputDto model.CreateSystemDictDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	var createDto = model.CreateSystemDictionaryDto{
		SystemDictionaryBase: model.SystemDictionaryBase{
			DictCode:  inputDto.DictCode,
			DictKey:   "ParentDict",
			DictValue: inputDto.DictValue,
			Remark:    inputDto.Remark,
			Sort:      0,
		},
	}

	if result, err := domainservice.SystemDictionaryDomainService.Insert(createDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Create doc
// @Summary 创建字典子集
// @Tags SystemDictionary
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemDictionaryDto true " "
// @Success 200 {object} model.SystemDictionary " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/create [POST]
func (app *systemDictionaryAppService) Create(c *gin.Context) {
	var inputDto model.CreateSystemDictionaryDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemDictionaryDomainService.Insert(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改字典
// @Tags SystemDictionary
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemDictionaryDto true " "
// @Success 200 {object} model.SystemDictionary " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/update [PUT]
func (app *systemDictionaryAppService) Update(c *gin.Context) {
	var inputDto model.UpdateSystemDictionaryDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemDictionaryDomainService.Update(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// UpdateStatus doc
// @Summary 修改字典状态
// @Tags SystemDictionary
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemDictionaryDto true " "
// @Success 200 {object} model.SystemDictionary " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/updateStatus [PUT]
func (app *systemDictionaryAppService) UpdateStatus(c *gin.Context) {
	var id = c.Query("id")
	if err := domainservice.SystemDictionaryDomainService.UpdateStatus(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// Delete doc
// @Summary 删除字典
// @Tags SystemDictionary
// @Description
// @Accept mpfd
// @Produce  json
// @Param id query string true " "
// @Success 200 {object} bool " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/delete [DELETE]
func (app *systemDictionaryAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	if err := domainservice.SystemDictionaryDomainService.Delete(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// GetList doc
// @Summary 查询字典列表
// @Tags SystemDictionary
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.SystemUserListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/list [GET]
func (app *systemDictionaryAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	dbSession := orm.QSession(false).Omit("Operation", "StatusText")
	if err := DBCondition(inputDto, dbSession, "", model.SystemDictionaryListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}

	var result = make([]model.SystemDictionaryListDto, 0)
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

	for i := 0; i < len(result); i++ {
		cDict := result[i]
		result[i].StatusText = domainconsts.GetConstFlag(result[i].Status, domainconsts.CommonStatusConstFlags)
		operations := make([]dto.OperationDto, 0)
		if cDict.DictKey == "ParentDict" {
			operations = append(operations, dto.GetOperationButton("DictConfig", "字典配置", "Infrastructure.SystemDictionary:Operation.DictConfig"))
		}
		operations = append(operations, dto.GetOperationButton("Modify", "修改", "Infrastructure.SystemDictionary:Operation.Modify"))
		if cDict.Status == domainconsts.SystemUserStatusDisable {
			operations = append(operations, dto.GetOperationButton("DisOrEnable", "启用", "Infrastructure.SystemDictionary:Operation.DisOrEnable"))
		} else {
			operations = append(operations, dto.GetOperationButton("DisOrEnable", "停用", "Infrastructure.SystemDictionary:Operation.DisOrEnable"))
		}
		operations = append(operations, dto.GetOperationButton("Delete", "删除", "Infrastructure.SystemDictionary:Operation.Delete"))
		result[i].Operation = GetCurrentUserGrantedOperation(operations)
	}
	resultDto.Data = result
	app.ResponseSuccess(c, resultDto)
}

// RefreshSystemDictCache doc
// @Summary 刷新缓存
// @Tags SystemDictionary
// @Description
// @Success 200 {object} bool " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/systemDictionary/refresh [POST]
func (app *systemDictionaryAppService) RefreshSystemDictCache(c *gin.Context) {
	if err := domainservice.SystemDictionaryDomainService.RefreshSystemDictCache(); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}
