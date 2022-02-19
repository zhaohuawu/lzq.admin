package domainservice

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/05
 * @Version 1.0.0
 */

import (
	"github.com/goinggo/mapstructure"
	"github.com/pkg/errors"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
)

// 租户领域服务
type tenantDomainService struct{}

var TenantDomainService = tenantDomainService{}

func (t *tenantDomainService) Insert(modelDto model.CreateTenantDto) (result model.Tenant, err error) {
	var entity model.Tenant
	var exist bool
	exist, err = orm.QSession(false).Exist(&model.Tenant{Code: modelDto.Code})
	if err != nil {
		return entity, err
	}
	if exist {
		return entity, errors.New("租户编码：" + modelDto.Code + " 已存在，请更换租户编码")
	}

	resultMap := utility.StructToMap(modelDto, false)
	if err = mapstructure.Decode(resultMap, &entity); err != nil {
		return entity, err
	}
	entity.ID = utility.UuidCreate()
	entity.Status = domainconsts.TenantStatusEnable

	_, err = orm.ISession().Insert(&entity)
	return entity, err
}
