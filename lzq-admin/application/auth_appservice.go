package application

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/dto"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/auth"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/tenant"
	"lzq-admin/pkg/utility"
	"strconv"
	"time"
)

type AuthAppService struct {
	BaseAppService
}

var IAuthAppService = AuthAppService{}

// GetCaptcha doc
// @Summary 获取验证码
// @Tags Auth
// @Description
// @Produce  json
// @Success 200 {object} dto.CaptchaDto " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/auth/captcha [GET]
func (s *AuthAppService) GetCaptcha(c *gin.Context) {
	captchaUrl, resultNum := utility.CaptchaUtil.GenerateCaptchaHandler()
	uuid := utility.UuidCreate()
	captchaKey := fmt.Sprintf("Captcha:%v", uuid)
	exp := time.Duration(60+utility.RandomNum(0, 10)) * time.Second
	cache.RedisUtil.NewRedis(false).Set(captchaKey, strconv.Itoa(resultNum), exp)
	s.ResponseSuccess(c, &dto.CaptchaDto{Key: uuid, CaptchaUrl: captchaUrl})
}

// Login doc
// @Summary 登录
// @Tags Auth
// @Description 登录服务
// @Accept json
// @Produce  json
// @Param object body dto.LoginDto true "登录参数"
// @Success 200 {object} dto.LoginTokenResponseDto
// @Failure 500 {object} ResponseDto
// @Router /api/app/auth/login [POST]
func (s *AuthAppService) Login(c *gin.Context) {
	responseDto := &dto.LoginTokenResponseDto{
		TokenType: "Bearer",
	}
	var inputDto dto.LoginDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		responseDto.IsError = true
		responseDto.ErrorDescription = err.Error()
		s.ResponseSuccess(c, responseDto)
		return
	}
	captchaKey := fmt.Sprintf("Captcha:%v", inputDto.CaptchaKey)
	captchaValue := cache.RedisUtil.NewRedis(false).Get(captchaKey)
	if len(captchaValue) == 0 {
		responseDto.IsError = true
		responseDto.ErrorDescription = "验证码已过期"
		s.ResponseSuccess(c, responseDto)
		return
	}
	if captchaValue != inputDto.CaptchaValue {
		responseDto.IsError = true
		responseDto.ErrorDescription = "验证码不正确"
		s.ResponseSuccess(c, responseDto)
		return
	}

	// 校验租户编码
	tenant, err := tenant.GetTenantByCode(inputDto.TenantCode)
	if err != nil {
		responseDto.IsError = true
		responseDto.Exception = err
		s.ResponseSuccess(c, responseDto)
		return
	}
	if tenant.Status != domainconsts.TenantStatusEnable {
		responseDto.IsError = true
		responseDto.ErrorDescription = fmt.Sprintf("租户已%v", domainconsts.GetConstFlag(tenant.Status, domainconsts.TenantConstFlags))
		s.ResponseSuccess(c, responseDto)
		return
	}

	var user model.SystemUser
	has, err2 := orm.DB.Where("LoginName=?", inputDto.LoginName).Where("TenantId=?", tenant.TenantId).And("IsDeleted=?", 0).Get(&user)
	if err2 != nil {
		responseDto.IsError = true
		responseDto.ErrorDescription = err2.Error()
		s.ResponseSuccess(c, responseDto)
		return
	}
	if has == false {
		responseDto.IsError = true
		responseDto.ErrorDescription = "账号不存在"
		s.ResponseSuccess(c, responseDto)
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputDto.Password)); err != nil {
		responseDto.IsError = true
		responseDto.Exception = errors.New("密码不正确")
		s.ResponseSuccess(c, responseDto)
		return
	}
	responseDto.ID = user.ID
	responseDto.Name = user.LoginName
	if _, err3 := utility.StringToUuid(tenant.TenantId); err3 != nil {
		responseDto.IsError = true
		responseDto.ErrorDescription = err3.Error()
		s.ResponseSuccess(c, responseDto)
		return
	}
	accessToken, err4 := token.GenerateToken(user.ID, user.LoginName, user.UserName, token.SysTypeAdmin, tenant.TenantId)
	if err4 != nil {
		responseDto.IsError = true
		responseDto.ErrorDescription = err4.Error()
		s.ResponseSuccess(c, responseDto)
		return
	}
	responseDto.AccessToken = accessToken
	s.ResponseSuccess(c, responseDto)
	//fmt.Sprintf("Bearer %v", accessToken))
}

// Logout doc
// @Summary 退出登录
// @Tags Auth
// @Description
// @Produce  json
// @Success 200 {object} bool true：退出成功
// @Failure 500 {object} ResponseDto
// @Router /api/app/auth/logOut [POST]
func (s *AuthAppService) Logout(c *gin.Context) {
	s.ResponseSuccess(c, true)
}
