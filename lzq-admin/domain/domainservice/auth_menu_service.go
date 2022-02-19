package domainservice

import (
	"github.com/pkg/errors"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/middleware"
	"lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"sync"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/2
 * @Version 1.0.0
 */

// AuthMenuDomainService 菜单领域服务
type authMenuDomainService struct {
	wg sync.WaitGroup
}

var AuthMenuDomainService = authMenuDomainService{}

func (s *authMenuDomainService) Insert(modelDto model.CreateAuthMenuDto) (model.AuthMenu, error) {
	var entity model.AuthMenu
	entity = model.AuthMenu{
		AuthMenuBase: model.AuthMenuBase{
			ModuleId: modelDto.ModuleId,
			Code:     modelDto.Code,
			Name:     modelDto.Name,
			Rank:     modelDto.Rank,
			Policy:   modelDto.Policy,
			Icon:     modelDto.Icon,
			ParentId: modelDto.ParentId,
			IsBranch: modelDto.IsBranch,
			IsHidden: modelDto.IsHidden,
			Path:     modelDto.Path,
			Url:      modelDto.Url,
		},
	}
	entity.ID = utility.UuidCreate()
	var dbSession, err = orm.BeginTrans()
	if err != nil {
		hsflogger.LogError("", err)
	}
	if modelDto.IsBranch == false {
		_, err = AuthPermissionDomainService.Insert(dbSession, model.CreateAuthPermissionDto{
			AuthPermissionBase: model.AuthPermissionBase{
				MenuId:          entity.ID,
				Name:            "页面访问",
				Code:            "Access",
				Rank:            1,
				Policy:          "View.Access",
				PermissionGroup: domainconsts.PermissionGroupView,
			},
		})
	}

	_, err = orm.ISessionWithTrans(dbSession).Insert(&entity)

	if err != nil {
		dbSession.Rollback()
		return model.AuthMenu{}, err
	}
	err = dbSession.Commit()
	return entity, err
}

func (s *authMenuDomainService) Get(id string) (model.AuthMenuDto, error) {
	var dto model.AuthMenuDto
	has, err := orm.QSession(false).ID(id).Get(&dto)
	if err != nil {
		return dto, err
	}
	if has {
		return dto, nil
	} else {
		return dto, errors.New("菜单不存在或已删除")
	}
}

func (s *authMenuDomainService) Update(inputDto model.UpdateAuthMenuDto) (model.AuthMenu, error) {
	var m model.AuthMenu
	var has int64
	var err error
	if has, err = orm.QSession(false).ID(inputDto.Id).Count(&m); err != nil {
		return model.AuthMenu{}, err
	}
	if has <= 0 {
		return model.AuthMenu{}, errors.New("菜单不存在")
	}
	m.Name = inputDto.Name
	m.Code = inputDto.Code
	m.Rank = inputDto.Rank
	m.IsHidden = inputDto.IsHidden
	m.Path = inputDto.Path
	m.Url = inputDto.Url
	m.IsBranch = inputDto.IsBranch
	m.ParentId = inputDto.ParentId
	m.Icon = inputDto.Icon
	m.ModuleId = inputDto.ModuleId
	m.Policy = inputDto.Policy

	var updateFields []string
	updateFields, err = orm.GetUpdateFields(model.UpdateAuthMenuDto{})
	if err != nil {
		return model.AuthMenu{}, err
	}

	var updateNum int64
	if updateNum, err = orm.USession(false).Cols(updateFields...).ID(inputDto.Id).Update(&m); err != nil {
		return model.AuthMenu{}, err
	}
	if updateNum < 1 {
		return model.AuthMenu{}, errors.New("修改失败")
	}
	// TODO 容器菜单改成非容器菜单，需要往AuthPermission表中插入页面访问操作按钮
	return m, nil
}

func (s *authMenuDomainService) Delete(id string) error {
	m := model.AuthMenu{
		BaseModel: model.BaseModel{
			IsDeleted:    true,
			DeleterId:    middleware.TokenClaims.Id,
			DeletionTime: time.Now(),
		},
	}
	dbSession, err := orm.BeginTrans()
	if err != nil {
		return err
	}
	s.wg.Add(1)
	var err1 error
	_, err = orm.DSessionWithTrans(false, dbSession).ID(id).Update(&m)
	go func() {
		err1 = AuthPermissionDomainService.Delete(dbSession, id)
		s.wg.Done()
	}()
	s.wg.Wait()
	if err != nil || err1 != nil {
		dbSession.Rollback()
		if err1 != nil {
			err = err1
		}
		return err
	}
	err = dbSession.Commit()
	if err != nil {
		return err
	}
	return nil
}
