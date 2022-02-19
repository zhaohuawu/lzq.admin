package application

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/model"
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
	//userId := middleware.TokenClaims.Id
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

	result := grantedMenuTree(pMenus, menus)

	app.ResponseSuccess(c, result)
}
func grantedMenuTree(parentMenus []model.UserGrantedMenuDto, menus []model.UserGrantedMenuDto) []model.UserGrantedMenuDto {
	// 排序 -- 升序 由小到大  使用大于号>表示降序，小于号<表示升序
	sort.SliceStable(parentMenus, func(i int, j int) bool {
		return parentMenus[i].Rank < parentMenus[j].Rank
	})
	for i := 0; i < len(parentMenus); i++ {
		cMenus := make([]model.UserGrantedMenuDto, 0)
		linq.From(menus).WhereT(func(s model.UserGrantedMenuDto) bool {
			return s.ParentId == parentMenus[i].Id
		}).ToSlice(&cMenus)
		if len(cMenus) > 0 {
			parentMenus[i].Children = grantedMenuTree(cMenus, menus)
		}
	}
	return parentMenus
}
