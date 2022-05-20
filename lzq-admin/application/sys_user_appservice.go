package application

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

import (
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"lzq-admin/config"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/domain/model/extrastruct"
	token "lzq-admin/pkg/auth"
	"lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"sync"
)

type systemUserAppService struct {
	BaseAppService
	wg sync.WaitGroup
}

var ISystemUserAppService = systemUserAppService{}

// Create doc
// @Summary 创建系统用户接口
// @Tags SystemUser
// @Description
// @Accept json
// @Produce  json
// @Param object body model.CreateSystemUserDto true " "
// @Success 200 {object} model.SystemUserDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/sysUser [POST]
func (app *systemUserAppService) Create(c *gin.Context) {
	var inputDto model.CreateSystemUserDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	result, rerr := domainservice.SystemUserDomainService.Insert(inputDto)
	if rerr != nil {
		// gin.H封装了生成json数据的工具
		app.ResponseError(c, rerr)
		return
	}
	app.ResponseSuccess(c, result)
}

// Get doc
// @Summary 根据用户ID或者登陆名查询用户信息
// @Tags SystemUser
// @Description id和loginName必须传一个值
// @Accept mpfd
// @Produce  json
// @Param id query string false "用户ID"
// @Param loginName query string false "用户登录名"
// @Success 200 {object} model.SystemUserDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/get [GET]
func (app *systemUserAppService) Get(c *gin.Context) {
	id := c.Query("id")
	loginName := c.Query("loginName")

	var user model.SystemUser
	if err := domainservice.SystemUserDomainService.Get(&user, id, loginName); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		result := &model.SystemUserInfoDto{ID: user.ID,
			UserName:   user.UserName,
			Sex:        user.Sex,
			Status:     user.Status,
			LoginName:  user.LoginName,
			HeadImgURL: user.HeadImgURL,
			Mobile:     user.Mobile,
			Email:      user.Email,
		}
		result.RoleIds, err = domainservice.AuthCheckerDomainService.GetUserGrantedRoleIds(user.ID)
		if err != nil {
			app.ResponseError(c, err)
			return
		}
		app.ResponseSuccess(c, result)
	}
}

// GetCurrentUserInfo doc
// @Summary 查询当前用户（无缓存）
// @Tags SystemUser
// @Description
// @Accept mpfd
// @Produce  json
// @Success 200 {object} model.SystemUserInfoDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/currentUserInfo [GET]
func (app *systemUserAppService) GetCurrentUserInfo(c *gin.Context) {
	id := token.GlobalTokenClaims.Id
	var user model.SystemUser
	if err := domainservice.SystemUserDomainService.Get(&user, id, ""); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		result := &model.SystemUserInfoDto{
			ID:         user.ID,
			UserName:   user.UserName,
			Sex:        user.Sex,
			Status:     user.Status,
			LoginName:  user.LoginName,
			HeadImgURL: user.HeadImgURL,
			Mobile:     user.Mobile,
			Email:      user.Email,
		}
		app.ResponseSuccess(c, result)
	}
}

// Delete
// @Summary 删除用户
// @Tags SystemUser
// @Description
// @Produce  json
// @Param id query string true "用户ID"
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/user [DELETE]
func (app *systemUserAppService) Delete(c *gin.Context) {
	id := c.Query("id")
	var user model.SystemUser
	if _, err := orm.DSession(true).ID(id).Update(&user); err != nil {
		app.ResponseError(c, err)
		return
	}
	// 清除用户详情缓存
	domainservice.SystemUserDomainService.RemoveUserInfoById(id)
	app.ResponseSuccess(c, true)
}

// GetList doc
// @Summary 查询用户列表
// @Tags SystemUser
// @Description
// @Accept mpfd
// @Produce  json
// @Param object query PageParamsDto true " "
// @Success 200 {array} model.SystemUserListDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/sysUserList [GET]
func (app *systemUserAppService) GetList(c *gin.Context) {
	var inputDto PageParamsDto
	if err := c.ShouldBind(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	dbSession := orm.QSession(true, "u").Table(model.TableSystemUser).Alias("u").
		// Join("Left", model.TableAuthUserdataprivilege+" as urp", fmt.Sprintf("u.Id = urp.UserId and urp.IsDeleted=0 and urp.TenantId='%v'", token.GlobalTokenClaims.TenantId)).
		//Select("u.*,urp.RoleId").
		Omit("Operation", "StatusName", "HeadImgLink", "Password", "RoleIds")
	if err := DBCondition(inputDto, dbSession, "u", model.SystemUserListDto{}); err != nil {
		app.ResponseError(c, err)
		return
	}
	var result = make([]model.SystemUserListDto, 0)
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
		result[i].StatusName = domainconsts.GetConstFlag(result[i].Status, domainconsts.SystemUserConstFlags)
		operations := make([]dto.OperationDto, 0)
		operations = append(operations, dto.GetOperationButton("Edit", "编辑", "Infrastructure.SysUser:Operation.Edit"))
		if result[i].Status == domainconsts.SystemUserStatusDisable {
			operations = append(operations, dto.GetOperationButton("EditStatus", "启用", "Infrastructure.SysUser:Operation.EditStatus"))
		} else {
			operations = append(operations, dto.GetOperationButton("EditStatus", "停用", "Infrastructure.SysUser:Operation.EditStatus"))
		}
		operations = append(operations, dto.GetOperationButton("Delete", "删除", "Infrastructure.SysUser:Operation.Delete"))
		operations = append(operations, dto.GetOperationButton("UpdatePassword", "修改密码", "Infrastructure.SysUser:Operation.UpdatePassword"))
		result[i].Operation = GetCurrentUserGrantedOperation(operations)
	}
	resultDto.Data = result
	app.ResponseSuccess(c, resultDto)
}

// Update doc
// @Summary 修改用户
// @Tags SystemUser
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemUserDto true " "
// @Success 200 {object} ResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/editSysUser [POST]
func (app *systemUserAppService) Update(c *gin.Context) {
	var inputDto model.UpdateSystemUserDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	if _, err := domainservice.SystemUserDomainService.Update(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	app.ResponseSuccess(c, true)
}

// GetUserInfo doc
// @Summary 查询当前用户信息
// @Tags SystemUser
// @Description
// @Produce  json
// @Success 200 {object} model.SystemUserInfoDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/userInfo [GET]
func (app *systemUserAppService) GetUserInfo(c *gin.Context) {
	userId := token.GlobalTokenClaims.Id
	if userInfo, err := domainservice.SystemUserDomainService.GetUserInfo(userId); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, userInfo)
	}
}

// UpdateSystemStatus doc
// @Summary 启用/停用用户
// @Tags SystemUser
// @Description
// @Accept json
// @Produce  json
// @Param object body dto.BaseDto true " "
// @Success 200 {object} bool true：修改成功
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/sysUserStatus [PUT]
func (app *systemUserAppService) UpdateSystemStatus(c *gin.Context) {
	var inputDto dto.BaseDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var user model.SystemUser
	if has, err := orm.QSession(true).ID(inputDto.ID).Get(&user); err != nil {
		app.ResponseError(c, err)
		return
	} else if !has {
		app.ResponseError(c, errors.New("用户不存在"))
		return
	}
	if user.Status == domainconsts.SystemUserStatusEnable {
		user.Status = domainconsts.SystemUserStatusDisable
	} else {
		user.Status = domainconsts.SystemUserStatusEnable
	}
	if effect, err := orm.USession(true).Cols("Status").ID(inputDto.ID).Update(&user); err != nil {
		app.ResponseError(c, err)
		return
	} else if effect <= 0 {
		app.ResponseError(c, errors.New("修改失败"))
		return
	} else {
		// 清除用户详情缓存
		domainservice.SystemUserDomainService.RemoveUserInfoById(inputDto.ID)
		app.ResponseSuccess(c, true)
	}
}

// UpdateSystemUserPassword doc
// @Summary 修改用户密码
// @Tags SystemUser
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemUserPasswordDto true " "
// @Success 200 object bool true：修改成功
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/editUserPassword [POST]
func (app *systemUserAppService) UpdateSystemUserPassword(c *gin.Context) {
	var inputDto model.UpdateSystemUserPasswordDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}

	if err := domainservice.SystemUserDomainService.UpdateSystemUserPassword(inputDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// UpdateCurrentUserPassword doc
// @Summary 修改当前用户密码
// @Tags SystemUser
// @Description
// @Accept json
// @Produce  json
// @Param object body model.UpdateSystemUserPasswordBaseDto true " "
// @Success 200 object bool true：修改成功
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/updateCurrentUserPassword [POST]
func (app *systemUserAppService) UpdateCurrentUserPassword(c *gin.Context) {
	var inputDto model.UpdateSystemUserPasswordBaseDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		app.ResponseError(c, err)
		return
	}
	var passwordDto = model.UpdateSystemUserPasswordDto{
		ID:                              token.GlobalTokenClaims.Id,
		UpdateSystemUserPasswordBaseDto: inputDto,
	}
	if err := domainservice.SystemUserDomainService.UpdateSystemUserPassword(passwordDto); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		app.ResponseSuccess(c, true)
	}
}

// GetDefaultAvatar doc
// @Summary 获取默认头像
// @Tags SystemUser
// @Description
// @Produce  json
// @Success 200 {string} string 默认头像
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysUser/defaultAvatar [GET]
func (app *systemUserAppService) GetDefaultAvatar(c *gin.Context) {
	hsflogger.LogInformation("测试日志")
	if sysConfig, err := domainservice.SysConfigDomainService.GetSysConfigCacheByCode(model.ExtraQiNiuConfig, "QiNiuStorage"); err != nil {
		app.ResponseError(c, err)
		return
	} else {
		var qnconfig extrastruct.ExtraQiNiuConfig
		if err := jsoniter.UnmarshalFromString(sysConfig, &qnconfig); err != nil {
			app.ResponseError(c, err)
			return
		}
		fileUrl := utility.UrlJoint(qnconfig.BaseUrl, qnconfig.Directory, config.Config.DefaultAvatar)
		app.ResponseSuccess(c, fileUrl)
	}
}
