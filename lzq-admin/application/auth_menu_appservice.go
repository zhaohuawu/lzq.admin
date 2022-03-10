package application

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"sort"
	"strings"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/1
 * @Version 1.0.0
 */

type authMenuAppService struct {
	BaseAppService
	wg sync.WaitGroup
}

var IAuthMenuAppService = authMenuAppService{}

// Create doc
// @Summary 创建菜单
// @Tags AuthMenu
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateAuthMenuDto true " "
// @Success 200 {object} model.AuthMenuDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/menu [POST]
func (app *authMenuAppService) Create(c *gin.Context) {
	var inputDto model.CreateAuthMenuDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	result, err := domainservice.AuthMenuDomainService.Insert(inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, result)
}

// Get doc
// @Summary 根据ID查询菜单
// @Tags AuthMenu
// @Description id
// @Accept mpfd
// @Produce  json
// @Param id query string true "菜单ID"
// @Success 200 {object} model.AuthMenuDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/get [GET]
func (app *authMenuAppService) Get(c *gin.Context) {
	id := c.Query("id")
	if result, err := domainservice.AuthMenuDomainService.Get(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, result)
	}
}

// Update doc
// @Summary 修改菜单
// @Tags AuthMenu
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateAuthMenuDto true " "
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/menu [PUT]
func (app *authMenuAppService) Update(c *gin.Context) {
	var inputDto model.UpdateAuthMenuDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	if _, err := domainservice.AuthMenuDomainService.Update(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// Delete
// @Summary 删除菜单
// @Tags AuthMenu
// @Description 删除菜单，同时删除菜单下的操作权限
// @Produce  json
// @Param id query string true "菜单ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/menu [DELETE]
func (app *authMenuAppService) Delete(c *gin.Context) {
	id := c.Query("MenuId")
	if err := domainservice.AuthMenuDomainService.Delete(id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// GetList doc
// @Summary 菜单列表
// @Tags AuthMenu
// @Description
// @Accept mpfd
// @Produce  json
// @Param object formData PageParamsDto true " "
// @Success 200 {array} model.AuthMenuListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/menusList [GET]
func (app *authMenuAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	dbSession := orm.QSession(false, "menu").Table(&model.AuthMenu{}).Alias("menu").
		Select("menu.*,module.Code as ModuleCode,module.Name as ModuleName").
		Join("Inner", model.TableAuthModule+" as module", "menu.ModuleId = module.Id").
		Where("module.IsDeleted = ?", 0)
	if err := DBCondition(inputDto, dbSession, "menu", model.AuthMenuListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result = make([]model.AuthMenuListDto, 0)
	var resultDto PageListDto
	var err error
	dbSession.Omit("Operation", "Children")
	if inputDto.RequireTotalCount {
		resultDto.TotalCount, err = dbSession.FindAndCount(&result)
	} else {
		err = dbSession.Find(&result)
	}
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	operation := GetCurrentUserGrantedOperation([]dto.OperationDto{
		dto.GetOperationButton("Edit", "编辑", "SystemSetup.Menu:Operation.Edit"),
		dto.GetOperationButton("Delete", "删除", "SystemSetup.Menu:Operation.Delete"),
	})
	// 整理出parentId=null的根菜单和子菜单
	var pMenus = make([]model.AuthMenuListDto, 0)
	var cMenus = make([]model.AuthMenuListDto, 0)
	for _, v := range result {
		v.Operation = operation
		if len(v.ParentId) == 0 {
			pMenus = append(pMenus, v)
		} else {
			cMenus = append(cMenus, v)
		}
	}
	if len(inputDto.Filter) > 0 {
		pMenus = append(pMenus, cMenus...)
		resultDto.Data = pMenus
	} else {
		resultDto.Data = menuTree(pMenus, cMenus)
	}
	app.ResponseSuccess(c, resultDto)
}
func menuTree(parentMenus []model.AuthMenuListDto, menus []model.AuthMenuListDto) []model.AuthMenuListDto {
	// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
	sort.SliceStable(parentMenus, func(i int, j int) bool {
		return parentMenus[i].Rank < parentMenus[j].Rank
	})
	for i := 0; i < len(parentMenus); i++ {
		cMenus := make([]model.AuthMenuListDto, 0)
		linq.From(menus).WhereT(func(s model.AuthMenuListDto) bool {
			return s.ParentId == parentMenus[i].ID
		}).ToSlice(&cMenus)
		if len(cMenus) > 0 {
			parentMenus[i].Children = menuTree(cMenus, menus)
		} else {
			parentMenus[i].Children = []model.AuthMenuListDto{}
		}
	}
	return parentMenus
}

// GetMenuList doc
// @Summary 根据菜单类型查询菜单
// @Tags AuthMenu
// @Description id
// @Accept mpfd
// @Produce  json
// @Param menuType query string true "菜单类型"
// @Success 200 {array} model.AuthMenuDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/menu/menuList [GET]
func (app *authMenuAppService) GetMenuList(c *gin.Context) {
	menuType := c.Query("menuType")

	var result = make([]model.AuthMenuDto, 0)
	dbSession := orm.QSession(false)
	if strings.ToLower(menuType) == "branch" {
		dbSession.Where("IsBranch = ?", true)
	} else if strings.ToLower(menuType) == "menu" {
		dbSession.Where("IsBranch = ?", false)
	}
	if err := dbSession.Find(&result); err != nil {
		app.ResponseError(c, err)
	} else {
		app.ResponseSuccess(c, result)
	}
}
