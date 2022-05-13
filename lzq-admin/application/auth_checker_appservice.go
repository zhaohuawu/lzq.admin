package application

import (
	"encoding/json"
	"errors"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
	"lzq-admin/middleware"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"sort"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/1/30
 * @Version 1.0.0
 */

type authCheckerAppService struct {
	BaseAppService
}

var IAuthCheckerAppService = authCheckerAppService{}

// GetGrantedMenus doc
// @Summary 查询当前用户已授权的菜单
// @Tags AuthChecker
// @Description
// @Produce  json
// @Success 200 {array} model.UserGrantedMenuDto “ ”
// @Failure 500 {object} ResponseDto
// @Router /api/app/authenticateChecker/grantedMenus [GET]
func (app *authCheckerAppService) GetGrantedMenus(c *gin.Context) {
	userId := middleware.TokenClaims.Id
	result := make([]model.UserGrantedMenuDto, 0)

	cacheKey := cache.LzqCacheKeyHelper.GetUserGrantedMenusCacheKey(userId)
	cacheJson := cache.RedisUtil.NewRedis(true).Get(cacheKey)
	if cacheJson != "" {
		_ = json.Unmarshal([]byte(cacheJson), &result)
		app.ResponseSuccess(c, result)
		return
	}
	var menus = make([]model.UserGrantedMenuDto, 0)
	dbSession := orm.QSession(false, "menu").
		Table(model.TableAuthMenu).Alias("menu").
		Omit("Children").
		Where("menu.IsHidden = ?", false)
	if err := dbSession.Find(&menus); err != nil {
		app.ResponseError(c, err)
		return
	}
	pMenus := make([]model.UserGrantedMenuDto, 0)
	linq.From(menus).WhereT(func(s model.UserGrantedMenuDto) bool {
		return s.IsBranch && len(s.ParentId) == 0
	}).ToSlice(&pMenus)

	if isSuperAdmin, err := domainservice.SystemUserDomainService.IsSuperAdmin(userId); err != nil {
		app.ResponseError(c, err)
		return
	} else if isSuperAdmin {
		result = grantedMenuTree(pMenus, menus)
	} else {
		rightMenus := make([]model.UserGrantedMenuDto, 0)
		for _, m := range menus {
			if len(m.Policy) > 0 {
				if isGranted := domainservice.CurrentUserPermissionChecker.IsGranted(m.Policy); isGranted {
					rightMenus = append(rightMenus, m)
				}
			} else {
				rightMenus = append(rightMenus, m)
			}
		}
		result = grantedMenuTree(pMenus, rightMenus)
	}
	cache.RedisUtil.NewRedis(true).SetInterfaceArray(cacheKey, result, 0)
	app.ResponseSuccess(c, result)
}
func grantedMenuTree(parentMenus []model.UserGrantedMenuDto, menus []model.UserGrantedMenuDto) []model.UserGrantedMenuDto {
	// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
	sort.SliceStable(parentMenus, func(i int, j int) bool {
		return parentMenus[i].Rank < parentMenus[j].Rank
	})
	pMenus := make([]model.UserGrantedMenuDto, 0)
	for i := 0; i < len(parentMenus); i++ {
		cMenus := make([]model.UserGrantedMenuDto, 0)
		linq.From(menus).WhereT(func(s model.UserGrantedMenuDto) bool {
			return s.ParentId == parentMenus[i].Id
		}).ToSlice(&cMenus)
		if len(cMenus) > 0 {
			parentMenus[i].Children = grantedMenuTree(cMenus, menus)
		}

		if parentMenus[i].IsBranch && len(parentMenus[i].Children) > 0 {
			pMenus = append(pMenus, parentMenus[i])
		}
		if !parentMenus[i].IsBranch {
			pMenus = append(pMenus, parentMenus[i])
		}
	}
	return pMenus
}

// GetCurrentUserGrantedPermissions doc
// @Summary 查询当前用户所授权的操作权限
// @Tags AuthChecker
// @Description
// @Produce  json
// @Success 200 {array} string " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/permissionChecker/grantedPermissions [GET]
func (app *authCheckerAppService) GetCurrentUserGrantedPermissions(c *gin.Context) {
	if permissions, err := domainservice.AuthCheckerDomainService.GetUserGrantedPermissions(middleware.TokenClaims.Id); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, permissions)
	}
}
// DeleteUserRole doc
// @Summary 删除用户数据授权
// @Tags AuthChecker
// @Description
// @Produce  json
// @Param userDataPrivilegeId query string true "用户数据授权ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/authorize/userRole [DELETE]
func (app *authCheckerAppService) DeleteUserRole(c *gin.Context) {
	userDataPrivilegeId := c.Query("userDataPrivilegeId")
	var dataPrivilege model.AuthUserDataPrivilege
	if has, err := orm.QSession(true).ID(userDataPrivilegeId).Get(&dataPrivilege); err != nil {
		app.ResponseError(c, err)
		return
	} else if !has {
		app.ResponseError(c, errors.New("该角色没有授权给该用户"))
		return
	}
	if isSuperAdmin, err := domainservice.SystemUserDomainService.IsSuperAdmin(dataPrivilege.UserId); err != nil {
		app.ResponseError(c, err)
		return
	} else if isSuperAdmin {
		app.ResponseError(c, errors.New("管理员账号不能进行更换角色"))
		return
	}
	num, err := orm.DSession(true).ID(userDataPrivilegeId).Update(&model.AuthUserDataPrivilege{})
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	if num <= 0 {
		app.ResponseError(c, errors.New("删除失败"))
		return
	}
	domainservice.AuthPrivilegeCacheService.DeleteFunctionPrivilegeCache()
	domainservice.AuthPrivilegeCacheService.DeleteDataPrivilegeCache()
	app.ResponseSuccess(c, true)
}
