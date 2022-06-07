package application

import (
	"errors"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	token "lzq-admin/pkg/auth"
	"lzq-admin/pkg/orm"
	"sort"
	"strings"
	"sync"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/4
 * @Version 1.0.0
 */

type authRoleAppService struct {
	BaseAppService
	wg sync.WaitGroup
}

var IAuthRoleAppService = authRoleAppService{}

// Create doc
// @Summary 创建角色
// @Tags AuthRole
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateAuthRoleDto true " "
// @Success 200 {object} model.AuthRoleDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/role [POST]
func (app *authRoleAppService) Create(c *gin.Context) {
	var inputDto model.CreateAuthRoleDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	result, err := domainservice.AuthRoleDomainService.Insert(inputDto)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, result)
}

// Get doc
// @Summary 根据ID查询角色
// @Tags AuthRole
// @Description
// @Accept mpfd
// @Produce  json
// @Param id query string false "ID"
// @Success 200 {object} model.AuthRoleDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/get [GET]
func (app *authRoleAppService) Get(c *gin.Context) {
	id := c.Query("id")
	var result model.AuthRoleDto
	has, err := orm.QSession(true).ID(id).Get(&result)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	if has == false {
		app.ResponseError(c, errors.New("菜单不存在或已删除"))
		return
	}
	app.ResponseSuccess(c, result)
}

// Update doc
// @Summary 修改角色
// @Tags AuthRole
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateAuthRoleDto true " "
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/role [PUT]
func (app *authRoleAppService) Update(c *gin.Context) {
	var inputDto model.UpdateAuthRoleDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var m model.AuthRole
	m.Name = inputDto.Name
	m.Code = inputDto.Code
	m.Description = inputDto.Description
	m.RoleStatus = inputDto.RoleStatus
	if updateFields, err := orm.GetUpdateFields(model.UpdateAuthRoleDto{}); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		if updateNum, err := orm.USession(true).Cols(updateFields...).ID(inputDto.Id).Update(&m); err != nil {
			app.ResponseError(c, err)
			return
		} else if updateNum < 1 {
			app.ResponseError(c, errors.New("角色ID不存在,修改失败"))
			return
		}
	}
	app.ResponseSuccess(c, true)
}

// Delete
// @Summary 删除角色
// @Tags AuthRole
// @Description
// @Produce  json
// @Param id query string true "模块ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/role [DELETE]
func (app *authRoleAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	dbSession, err := orm.BeginTrans()
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	if _, err := orm.DSessionWithTrans(true, dbSession).ID(id).Update(&model.AuthRole{}); err != nil {
		dbSession.Rollback()
		app.ResponseError(c, err)
		return
	}
	if _, err := orm.DSessionWithTrans(true, dbSession).Where("RoleId=?", id).Update(&model.AuthRolePermission{}); err != nil {
		dbSession.Rollback()
		app.ResponseError(c, err)
		return
	}
	err = dbSession.Commit()
	if err != nil {
		dbSession.Rollback()
		app.ResponseError(c, err)
		return
	}
	domainservice.AuthPrivilegeCacheService.DeleteRoleGrantedPermissionsCache(id)
	app.ResponseSuccess(c, true)
}

// GetList doc
// @Summary 查询角色列表
// @Tags AuthRole
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.AuthRoleListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/roleList [GET]
func (app *authRoleAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	dbSession := orm.QSession(true)
	if err := DBCondition(inputDto, dbSession, "", model.AuthRoleListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result = make([]model.AuthRoleListDto, 0)
	var resultDto PageListDto
	var err error
	dbSession.Omit("Operation", "RoleStatusText", "HaveAuthorizeUser")
	if inputDto.RequireTotalCount {
		resultDto.TotalCount, err = dbSession.FindAndCount(&result)
	} else {
		err = dbSession.Find(&result)
	}
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	authorizeUsers := make([]model.AuthorizeUser, 0)
	if err := orm.QSession(true, "p", "u").Table(model.TableAuthUserdataprivilege).Alias("p").
		Join("INNER", model.TableSystemUser+" as u", "p.UserId=u.Id").
		Select("u.LoginName,u.UserName,u.IsTenantAdmin,p.Id as UserDataPrivilegeID,p.RoleId").Find(&authorizeUsers); err != nil {
		app.ResponseError(c, err)
		return
	}
	for i := 0; i < len(result); i++ {
		result[i].RoleStatusText = domainconsts.GetConstFlag(result[i].RoleStatus, domainconsts.RoleStatusConstFlags)
		haveAuthorizeUser := make([]model.AuthorizeUser, 0)
		linq.From(authorizeUsers).WhereT(func(s model.AuthorizeUser) bool {
			return s.RoleId == result[i].ID
		}).ToSlice(&haveAuthorizeUser)
		result[i].HaveAuthorizeUser = haveAuthorizeUser
		operations := make([]dto.OperationDto, 0)
		operations = append(operations, dto.GetOperationButton("FunctionAuthorization", "功能授权", "Infrastructure.Role:Operation.FunctionAuthorization"))
		if strings.Index(result[i].Name, "系统管理员") < 0 {
			operations = append(operations, dto.GetOperationButton("Edit", "编辑", "Infrastructure.Role:Operation.Edit"))
			if result[i].RoleStatus == domainconsts.SystemUserStatusDisable {
				operations = append(operations, dto.GetOperationButton("EditStatus", "启用", "Infrastructure.Role:Operation.EditStatus"))
			} else {
				operations = append(operations, dto.GetOperationButton("EditStatus", "停用", "Infrastructure.Role:Operation.EditStatus"))
			}
			operations = append(operations, dto.GetOperationButton("Delete", "删除", "Infrastructure.Role:Operation.Delete"))
		}
		result[i].Operation = GetCurrentUserGrantedOperation(operations)
	}

	resultDto.Data = result
	app.ResponseSuccess(c, resultDto)
}

// UpdateRoleStatus doc
// @Summary 启用/停用角色
// @Tags AuthRole
// @Description
// @Accept json
// @Produce  json
// @Param object body dto.BaseDto true " "
// @Success 200 {object} bool true：修改成功
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/roleStatus [PUT]
func (app *authRoleAppService) UpdateRoleStatus(c *gin.Context) {
	var inputDto dto.BaseDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var role model.AuthRole
	if has, err := orm.QSession(true).ID(inputDto.ID).Get(&role); err != nil {
		app.ResponseError(c, err)
		return
	} else if !has {
		app.ResponseError(c, errors.New("角色不存在"))
		return
	}
	if role.RoleStatus == domainconsts.RoleStatusEnable {
		role.RoleStatus = domainconsts.SystemUserStatusDisable
	} else {
		role.RoleStatus = domainconsts.RoleStatusEnable
	}
	if effect, err := orm.USession(true).Cols("RoleStatus").ID(inputDto.ID).Update(&role); err != nil {
		app.ResponseError(c, err)
		return
	} else if effect <= 0 {
		app.ResponseError(c, errors.New("修改失败"))
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// GetEnanleRoles doc
// @Summary 查询启用中的角色
// @Tags AuthRole
// @Description
// @Produce  json
// @Success 200 {array} model.AuthRoleSimpleDto " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/role/roles [GET]
func (app *authRoleAppService) GetEnanleRoles(c *gin.Context) {
	var result = make([]model.AuthRoleSimpleDto, 0)
	selectFields, err := orm.GetOptionFields(model.AuthRoleSimpleDto{})
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	err1 := orm.QSession(true).Table(model.TableAuthRole).
		Select(strings.Join(selectFields, ",")).
		Where("roleStatus = ?", domainconsts.RoleStatusEnable).Find(&result)
	if err1 != nil {
		app.ResponseError(c, err1)
		return
	}
	app.ResponseSuccess(c, result)
}

// GetRolePermissionDatas doc
// @Summary 获取角色授权的操作权限
// @Tags AuthRole
// @Description
// @Accept mpfd
// @Produce  json
// @Param roleId path string true "角色ID"
// @Success 200 {array} dto.RolePermissionTree
// @Failure 500 {object} ResponseDto
// @Router /api/app/authorize/rolePermissionDatas/{roleId} [GET]
func (app *authRoleAppService) GetRolePermissionDatas(c *gin.Context) {
	roleId := c.Param("roleId")
	var permissions = make([]dto.RolePermissionTree, 0)
	dbSession := orm.QSession(false, "p").Table(model.TableAuthPermission).Alias("p").
		Join("LEFT", model.TableAuthRolePermission+" as rp", "p.Id=rp.PermissionId and rp.RoleId=? and rp.IsDeleted=? and rp.TenantId=?", roleId, 0, token.GlobalTokenClaims.TenantId).
		Select("p.Id,p.MenuId as ParentId,'Permission' as Type,p.Name,p.Rank,IFNULL(rp.IsGranted ,false) as IsGranted,false as IsBranch")
	if err := dbSession.Find(&permissions); err != nil {
		app.ResponseError(c, err)
		return
	}

	var menus = make([]dto.RolePermissionTree, 0)
	mDBSession := orm.QSession(false).Table(model.TableAuthMenu).
		Select("Id,ParentId,'Menu' as Type,Name,`Rank`,false as IsGranted,IsBranch")
	if err := mDBSession.Find(&menus); err != nil {
		app.ResponseError(c, err)
		return
	}
	haveCheckedKeys := make([]string, 0)
	linq.From(permissions).WhereT(func(w dto.RolePermissionTree) bool { return w.IsGranted }).
		SelectT(func(w dto.RolePermissionTree) string { return w.ID }).ToSlice(&haveCheckedKeys)

	perissionIds := make([]string, 0)
	linq.From(permissions).SelectT(func(w dto.RolePermissionTree) string { return w.ID }).ToSlice(&perissionIds)
	result := dto.RolePermissionTreeDto{
		HaveCheckedKeys:    haveCheckedKeys,
		PerissionIds:       perissionIds,
		RolePermissionTree: rolePermissionTree("", menus, permissions),
	}

	app.ResponseSuccess(c, result)
}
func rolePermissionTree(parentId string, menus []dto.RolePermissionTree, permissions []dto.RolePermissionTree) []dto.RolePermissionTree {
	childrens := make([]dto.RolePermissionTree, 0)
	linq.From(menus).WhereT(func(w dto.RolePermissionTree) bool { return w.ParentID == parentId }).ToSlice(&childrens)
	// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
	sort.SliceStable(childrens, func(i int, j int) bool {
		return childrens[i].Rank < childrens[j].Rank
	})
	for i := 0; i < len(childrens); i++ {
		childrens[i].Children = rolePermissionTree(childrens[i].ID, menus, permissions)
		if !childrens[i].IsBranch {
			pChildrens := make([]dto.RolePermissionTree, 0)
			linq.From(permissions).WhereT(func(w dto.RolePermissionTree) bool { return w.ParentID == childrens[i].ID }).ToSlice(&pChildrens)
			// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
			sort.SliceStable(pChildrens, func(i int, j int) bool {
				return pChildrens[i].Rank < pChildrens[j].Rank
			})
			childrens[i].Children = append(childrens[i].Children, pChildrens...)
		}
	}
	return childrens
}

func (app *authRoleAppService) GrantPermissions(c *gin.Context) {
	var inputDto []model.CreateAuthRolePermissionDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if _, err := domainservice.AuthRolePermissionDomainService.Insert(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	app.wg.Add(2)
	go func() {
		domainservice.AuthPrivilegeCacheService.DeleteFunctionPrivilegeCache()
		app.wg.Done()
	}()
	go func() {
		roleIds := make([]string, 0)
		linq.From(inputDto).SelectT(func(w model.CreateAuthRolePermissionDto) string {
			return w.RoleId
		}).Distinct().ToSlice(&roleIds)
		domainservice.AuthPrivilegeCacheService.DeleteRoleGrantedPermissionsCache(roleIds...)
		app.wg.Done()
	}()
	app.wg.Wait()
	app.ResponseSuccess(c, true)
}
