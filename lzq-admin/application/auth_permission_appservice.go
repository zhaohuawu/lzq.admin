package application

import (
	"errors"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"sort"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/3
 * @Version 1.0.0
 */

type authPermissionAppService struct {
	BaseAppService
	wg sync.WaitGroup
}

var IAuthPermissionAppService = authPermissionAppService{}

// Create doc
// @Summary 创建操作权限
// @Tags AuthPermission
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateAuthPermissionDto true " "
// @Success 200 {object} model.AuthPermissionDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/permission [POST]
func (app *authPermissionAppService) Create(c *gin.Context) {
	var inputDto model.CreateAuthPermissionDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result model.AuthPermission
	dbSession, err := orm.BeginTrans()
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	result, err = domainservice.AuthPermissionDomainService.Insert(dbSession, inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	err = dbSession.Commit()
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, result)
}

// Get doc
// @Summary 根据ID查询
// @Tags AuthPermission
// @Description id
// @Accept mpfd
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} model.AuthPermissionDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/get [GET]
func (app *authPermissionAppService) Get(c *gin.Context) {
	id := c.Query("id")
	if result, err := domainservice.AuthPermissionDomainService.Get(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改操作权限
// @Tags AuthPermission
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateAuthPermissionDto true " "
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/permission [PUT]
func (app *authPermissionAppService) Update(c *gin.Context) {
	var inputDto model.UpdateAuthPermissionDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var m model.AuthPermission
	var has int64
	var err error
	if has, err = orm.QSession(false).ID(inputDto.Id).Count(&m); err != nil {
		app.ResponseError(c, err)
		return
	}
	if has <= 0 {
		app.ResponseError(c, errors.New("操作权限不存在"))
		return
	}
	m.Name = inputDto.Name
	m.Code = inputDto.Code
	m.Rank = inputDto.Rank
	m.Policy = inputDto.Policy
	m.MenuId = inputDto.MenuId

	var updateNum int64
	if updateNum, err = orm.USession(false).AllCols().ID(inputDto.Id).Update(&m); err != nil {
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
// @Summary 删除操作权限
// @Tags AuthPermission
// @Description
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/permission [DELETE]
func (app *authPermissionAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	m := model.AuthPermission{}
	_, err := orm.DSession(false).ID(id).Update(&m)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, true)
}

// GetList doc
// @Summary 操作权限列表
// @Tags AuthPermission
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.AuthPermissionListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/data [GET]
func (app *authPermissionAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	//操作权限
	dbSession := orm.QSession(false, "p", "m").Table(&model.AuthPermission{}).Alias("p").
		Select("p.*,'Permission' as TypeCode,p.MenuId as ParentId,m.Name as MenuName").
		Join("Inner", model.TableAuthMenu+" as m", "p.MenuId = m.Id")
	if err := DBCondition(inputDto, dbSession, "p", model.AuthPermissionListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var permissions = make([]model.AuthPermissionListDto, 0)
	var resultDto PageListDto
	var err error
	dbSession.Omit("Operation", "Children")
	err = dbSession.Asc("p.Rank").Find(&permissions)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	//菜单
	var menuIds = make([]string, 0)
	if len(inputDto.Filter) > 0 {
		linq.From(permissions).SelectT(func(s model.AuthPermissionListDto) string {
			return s.MenuId
		}).Distinct().ToSlice(&menuIds)
	}

	var menus = make([]model.AuthPermissionListDto, 0)
	mDBSession := orm.QSession(false).Table(&model.AuthMenu{}).Select("Id, ParentId, `Rank`, Name as MenuName, Policy,IsBranch, 'Menu' as TypeCode").Asc("Rank")
	if len(menuIds) > 0 {
		mDBSession.In("Id", menuIds)
	}
	err = mDBSession.Find(&menus)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	operations := GetCurrentUserGrantedOperation([]dto.OperationDto{
		dto.GetOperationButton("Edit", "编辑", "SystemSetup.Permission:Operation.Edit"),
		dto.GetOperationButton("Delete", "删除", "SystemSetup.Permission:Operation.Delete"),
	})
	// 整理出parentId=null的根菜单和子菜单
	var pMenus = make([]model.AuthPermissionListDto, 0)
	var cMenus = make([]model.AuthPermissionListDto, 0)
	var nPermissions = permissions
	for _, v := range menus {
		if v.IsBranch && len(v.ParentId) == 0 {
			pMenus = append(pMenus, v)
		} else {
			var npChildres = make([]model.AuthPermissionListDto, 0)
			var ipChildres = make([]model.AuthPermissionListDto, 0)
			if v.IsBranch == false {
				for _, pv := range nPermissions {
					pv.Children = make([]model.AuthPermissionListDto, 0)
					if v.ID == pv.MenuId {
						pv.ActualPolicy = v.Policy + ":" + pv.Policy
						pv.PermissionGroupText = domainconsts.GetConstFlag(pv.PermissionGroup, domainconsts.PermissionGroupConstFlags)
						pv.Operation = operations
						ipChildres = append(ipChildres, pv)
					} else {
						npChildres = append(npChildres, pv)
					}
				}
			}
			nPermissions = npChildres
			v.Children = ipChildres
			cMenus = append(cMenus, v)
		}
	}
	if len(inputDto.Filter) > 0 {
		resultDto.Data = cMenus
	} else {
		resultDto.Data = permissionTree(pMenus, cMenus)
	}
	app.ResponseSuccess(c, resultDto)
}
func permissionTree(parentMenus []model.AuthPermissionListDto, menus []model.AuthPermissionListDto) []model.AuthPermissionListDto {
	// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
	sort.SliceStable(parentMenus, func(i int, j int) bool {
		return parentMenus[i].Rank < parentMenus[j].Rank
	})
	for i := 0; i < len(parentMenus); i++ {
		cMenus := make([]model.AuthPermissionListDto, 0)
		linq.From(menus).WhereT(func(s model.AuthPermissionListDto) bool {
			return s.ParentId == parentMenus[i].ID
		}).ToSlice(&cMenus)
		if len(cMenus) > 0 {
			parentMenus[i].Children = permissionTree(cMenus, menus)
		}
	}
	return parentMenus
}

// GetPermissionGroup doc
// @Summary 获取所有的操作权限组
// @Tags AuthPermission
// @Description
// @Produce  json
//@Success 200 {array} dto.KeyAndValueDto “ ”
// @Failure 500 {object} ResponseDto
// @Router /api/app/permission/permissionGroup [GET]
func (app *authPermissionAppService) GetPermissionGroup(c *gin.Context) {
	result := make([]dto.KeyAndValueDto, 0)
	for k, v := range domainconsts.PermissionGroupConstFlags {
		result = append(result, dto.KeyAndValueDto{Key: k, Value: v})
	}
	app.ResponseSuccess(c, result)
}
