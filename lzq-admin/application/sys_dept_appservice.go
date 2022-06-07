package application

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/23
 * @Version 1.0.0
 */

type systemDeptAppService struct {
	BaseAppService
}

var SystemDeptAppService = systemDeptAppService{}

// Create doc
// @Summary 新增部门
// @Tags SystemDept
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemDeptDto true " "
// @Success 200 {object} model.SystemDept 其他描述
// @Failure 500 {object} ResponseDto
// @Router /api/app/dept/create [POST]
func (app *systemDeptAppService) Create(c *gin.Context) {
	var inputDto model.CreateSystemDeptDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemDeptDomainService.Insert(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改部门
// @Tags SystemDept
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemDeptDto true " "
// @Success 200 {object} model.SystemDept 其他描述
// @Failure 500 {object} ResponseDto
// @Router /api/app/dept/update [PUT]
func (app *systemDeptAppService) Update(c *gin.Context) {
	var inputDto model.UpdateSystemDeptDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if result, err := domainservice.SystemDeptDomainService.Update(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Delete
// @Summary 删除部门
// @Tags SystemDept
// @Description
// @Produce  json
// @Param id query string true "部门ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/dept/delete [DELETE]
func (app *systemDeptAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	if err := domainservice.SystemDeptDomainService.Delete(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// GetCompanyAndDeptList doc
// @Summary 公司和部门的树形列表
// @Tags SystemDept
// @Description
// @Accept json
// @Produce  json
// @Success 200 {array} model.SystemCompanyAndDeptListDto " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/dept/companyAndDeptList [GET]
func (app *systemDeptAppService) GetCompanyAndDeptList(c *gin.Context) {
	//var inputDto PageParamsDto
	//if err := c.ShouldBind(&inputDto); err != nil {
	//	app.ResponseError(c, err)
	//	return
	//}
	operation := GetCurrentUserGrantedOperation([]dto.OperationDto{
		dto.GetOperationButton("Edit", "编辑", "SystemSetup.Menu:Operation.Edit"),
		dto.GetOperationButton("Delete", "删除", "SystemSetup.Menu:Operation.Delete"),
	})
	cList := make([]model.SystemCompany, 0)
	if err := orm.QSession(true).Select("ID,Name,ParentId,Remark,`Rank`").OrderBy("`Rank`").Find(&cList); err != nil {
		app.ResponseError(c, err)
		return
	}
	dList := make([]model.SystemDept, 0)
	if err := orm.QSession(true).Select("ID,Name,CompanyId,ParentId,Remark,`Rank`").OrderBy("`Rank`").Find(&dList); err != nil {
		app.ResponseError(c, err)
		return
	}
	list := CompanyTree("", cList, dList, operation)
	app.ResponseSuccess(c, list)
}
func CompanyTree(parentId string, cList []model.SystemCompany, dList []model.SystemDept, operation string) []model.SystemCompanyAndDeptListDto {
	pList := make([]model.SystemCompany, 0)
	if len(parentId) == 0 {
		linq.From(cList).WhereT(func(s model.SystemCompany) bool {
			return len(s.ParentId) == 0
		}).ToSlice(&pList)
	} else {
		linq.From(cList).WhereT(func(s model.SystemCompany) bool {
			return s.ParentId == parentId
		}).ToSlice(&pList)
	}
	result := make([]model.SystemCompanyAndDeptListDto, 0)
	for _, c := range pList {
		list := make([]model.SystemCompanyAndDeptListDto, 0)
		l := CompanyTree(c.ID, cList, dList, operation)
		if len(l) > 0 {
			list = l
		}

		dChilds := make([]model.SystemDept, 0)
		linq.From(dList).WhereT(func(s model.SystemDept) bool {
			return s.CompanyId == c.ID && len(s.ParentId)==0
		}).ToSlice(&dChilds)
		if len(dChilds) > 0 {
			for _, dc := range dChilds {
				dr := model.SystemCompanyAndDeptListDto{}
				dr.ID = dc.ID
				dr.Type = "Dept"
				dr.Operation = operation
				dr.Name = dc.Name
				dr.CompanyId = dc.CompanyId
				dr.ParentId = dc.ParentId
				dr.Remark = dc.Remark
				dr.Rank = dc.Rank
				l := DeptTree(dc.ID, dList, operation)
				if len(l) > 0 {
					dr.Children = l
				}
				list = append(list, dr)
			}
		}
		d := model.SystemCompanyAndDeptListDto{}
		d.ID = c.ID
		d.Type = "Company"
		d.Operation = operation
		d.Name = c.Name
		d.ParentId = c.ParentId
		d.Remark = c.Remark
		d.Rank = c.Rank

		if len(list) > 0 {
			d.Children = list
		}
		d.Operation = operation
		result = append(result, d)
	}
	return result
}

func DeptTree(parentId string, dList []model.SystemDept, operation string) []model.SystemCompanyAndDeptListDto {
	pList := make([]model.SystemDept, 0)
	if len(parentId) == 0 {
		linq.From(dList).WhereT(func(s model.SystemDept) bool {
			return len(s.ParentId) == 0
		}).ToSlice(&pList)
	} else {
		linq.From(dList).WhereT(func(s model.SystemDept) bool {
			return s.ParentId == parentId
		}).ToSlice(&pList)
	}
	result := make([]model.SystemCompanyAndDeptListDto, 0)
	for _, c := range pList {
		d := model.SystemCompanyAndDeptListDto{}
		d.ID = c.ID
		d.Type = "Dept"
		d.Operation = operation
		d.Name = c.Name
		d.CompanyId = c.CompanyId
		d.ParentId = c.ParentId
		d.Remark = c.Remark
		d.Rank = c.Rank
		dChilds := DeptTree(c.ID, dList, operation)
		if len(dChilds) > 0 {
			d.Children = dChilds
		}

		result = append(result, d)
	}
	return result
}
