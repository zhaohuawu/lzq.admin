package domainservice

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/05
 * @Version 1.0.0
 */

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/cache"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

// 系统用户领域服务
type systemUserDomainService struct{}

var SystemUserDomainService = systemUserDomainService{}

func (u *systemUserDomainService) Insert(modelDto model.CreateSystemUserDto) (result model.SystemUser, err error) {
	var entity model.SystemUser
	var exist bool
	exist, err = orm.QSession(true).Where("LoginName=?", modelDto.LoginName).Exist(&model.SystemUser{})
	if err != nil {
		return entity, err
	}
	if exist {
		return entity, errors.New("登陆名：" + modelDto.LoginName + " 已存在，请更换登录名")
	}
	entity = model.SystemUser{
		SystemUserBase: model.SystemUserBase{
			LoginName:  modelDto.LoginName,
			UserName:   modelDto.UserName,
			Email:      modelDto.Email,
			HeadImgURL: modelDto.HeadImgURL,
			Sex:        modelDto.Sex,
			Mobile:     modelDto.Mobile,
		},
	}
	entity.ID = utility.UuidCreate()
	entity.Status = domainconsts.SystemUserStatusEnable

	if entity.Password != modelDto.SurePassword {
		return model.SystemUser{}, errors.New("两次输入的密码不同，请重新输入")
	}
	// 密码加密不可解析密码串
	phash, err1 := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
	if err1 != nil {
		return entity, err1
	}
	entity.Password = string(phash)
	dbSession, errT := orm.BeginTrans()
	if errT != nil {
		return model.SystemUser{}, err
	}

	if _, err := AuthUserdataPrivilegeDomainService.Insert(dbSession, entity.ID, modelDto.RoleIds); err != nil {
		dbSession.Rollback()
		return model.SystemUser{}, err
	}
	if _, err := orm.ISessionWithTrans(dbSession).Insert(&entity); err != nil {
		dbSession.Rollback()
		return model.SystemUser{}, err
	}
	if err := dbSession.Commit(); err != nil {
		dbSession.Rollback()
		return model.SystemUser{}, err
	}
	return entity, nil
}

func (u *systemUserDomainService) Update(inputDto model.UpdateSystemUserDto) (model.SystemUser, error) {
	var user model.SystemUser
	var has int64
	var err error
	if has, err = orm.QSession(true).ID(inputDto.ID).Count(&user); err != nil {
		return model.SystemUser{}, err
	}
	if has <= 0 {
		return model.SystemUser{}, errors.New("用户不存在")
	}

	var uEntity model.SystemUser
	uEntity.UserName = inputDto.UserName
	uEntity.Email = inputDto.Email
	uEntity.HeadImgURL = inputDto.HeadImgURL
	uEntity.Mobile = inputDto.Mobile
	uEntity.Sex = inputDto.Sex

	dbSession, errT := orm.BeginTrans()
	if errT != nil {
		return model.SystemUser{}, err
	}
	if len(inputDto.RoleIds) == 0 {
		if err := AuthUserdataPrivilegeDomainService.Delete(dbSession, inputDto.ID, ""); err != nil {
			dbSession.Rollback()
			return model.SystemUser{}, err
		}
	} else {
		if _, err := AuthUserdataPrivilegeDomainService.Insert(dbSession, inputDto.ID, inputDto.RoleIds); err != nil {
			dbSession.Rollback()
			return model.SystemUser{}, err
		}
	}

	if updateNum, err1 := orm.USessionWithTrans(true, dbSession).ID(inputDto.ID).Update(&uEntity); err1 != nil {
		return model.SystemUser{}, err1
	} else if updateNum < 1 {
		return model.SystemUser{}, errors.New("修改失败")
	}
	// 清除用户详情缓存
	u.RemoveUserInfoById(inputDto.ID)
	return uEntity, nil
}

// Get 根据用户ID或者LoginName查询用户信息
func (u *systemUserDomainService) Get(m *model.SystemUser, id, loginName string) error {
	var has bool = false
	var err error

	if id != "" {
		has, err = orm.QSession(true).ID(id).Get(m)
	} else if loginName != "" {
		has, err = orm.QSession(true).Where("LoginName=?", loginName).Get(m)
	}

	if err != nil {
		return err
	}
	if has {
		return nil
	} else {
		return errors.New("账号不存在")
	}
}

func (u *systemUserDomainService) GetUserInfo(userId string) (model.SystemUserInfoDto, error) {
	key := fmt.Sprintf("%v:%v", cache.LzqCacheHelper.GetCacheVersion(cache.LzqCacheTypeSysUser), userId)
	r := cache.RedisUtil.NewRedis(true, "UserInfo")
	userJson := r.Get(key)
	var userInfo model.SystemUserInfoDto
	if userJson != "" {
		if err := jsoniter.UnmarshalFromString(userJson, &userInfo);err!=nil{
			return userInfo,nil
		}
		return userInfo, nil
	}
	var user model.SystemUser
	if err := u.Get(&user, userId, ""); err != nil {
		return userInfo, err
	}
	userInfo = model.SystemUserInfoDto{
		RoleIds:       nil,
		RoleName:      "",
		Status:        user.Status,
		LoginName:     user.LoginName,
		UserName:      user.UserName,
		HeadImgURL:    user.HeadImgURL,
		HeadImgLink:   user.HeadImgURL,
		Sex:           user.Sex,
		Mobile:        user.Mobile,
		Email:         user.Email,
		IsTenantAdmin: user.IsTenantAdmin,
		ID:            user.ID,
	}

	v, isHave := user.ExtraProperties["SuperAdmin"]
	userInfo.SuperAdmin = isHave && v.(bool)

	r.SSet(key, userInfo, 0)
	return userInfo, nil
}

func (u *systemUserDomainService) IsTenantAdmin(userId string) (bool, error) {
	if userInfo, err := SystemUserDomainService.GetUserInfo(userId); err != nil {
		return false, err
	} else {
		return userInfo.IsTenantAdmin, nil
	}
}

func (u *systemUserDomainService) IsSuperAdmin(userId string) (bool, error) {
	if userInfo, err := SystemUserDomainService.GetUserInfo(userId); err != nil {
		return false, err
	} else {
		return userInfo.SuperAdmin, nil
	}
}

func (u *systemUserDomainService) RemoveUserInfoById(userId string) {
	key := fmt.Sprintf("%v:%v", cache.LzqCacheHelper.GetCacheVersion(cache.LzqCacheTypeSysUser), userId)
	cache.RedisUtil.NewRedis(true).Delete(key)
}
