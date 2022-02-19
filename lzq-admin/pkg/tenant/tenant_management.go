package tenant

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"errors"
	"fmt"
	"lzq-admin/config"
	"lzq-admin/config/appsettings"
	"strings"
)

func CheckUseMultiTenancy() (bool, appsettings.TenantInfo) {
	if config.Config.ServerConfig.UseMultiTenancy {
		return true, appsettings.TenantInfo{}
	} else {
		return false, appsettings.TenantInfo{
			TenantId: "11111111-1111-1111-1111-111111111111",
			Status:   "Enable",
			Name:     "默认租户",
			Code:     "Default",
		}
	}

}

// CheckTenantById 根据租户Id校验租户Id是否存在
func CheckTenantById(tenantId string) error {
	if u, _ := CheckUseMultiTenancy(); u == false {
		return nil
	}
	t, err := GetTenantById(tenantId)
	if err != nil {
		return err
	}
	fmt.Println("通过租户ID校验的租户信息：", t)
	return nil
}

// GetTenantById 根据租户ID获取租户信息
func GetTenantById(tenantId string) (tenant appsettings.TenantInfo, err error) {
	if u, d := CheckUseMultiTenancy(); u == false {
		return d, nil
	}
	for _, t := range appsettings.AppSettings.TenantInfos {
		if t.TenantId == tenantId {
			tenant = t
			return tenant, nil
		}
	}
	return tenant, errors.New("租户ID不存在")
}

// CheckTenantByCode 根据租户Code校验租户编码是否存在
func CheckTenantByCode(tenantCode string) error {
	if u, _ := CheckUseMultiTenancy(); u == false {
		return nil
	}
	t, err := GetTenantByCode(tenantCode)
	if err != nil {
		return err
	}
	fmt.Println("通过租户Code校验的租户信息：", t)
	return nil
}

// GetTenantByCode 根据租户Code获取租户信息
func GetTenantByCode(tenantCode string) (tenant appsettings.TenantInfo, err error) {
	if u, d := CheckUseMultiTenancy(); u == false {
		return d, nil
	}
	for _, t := range appsettings.AppSettings.TenantInfos {
		if strings.ToLower(t.Code) == strings.ToLower(tenantCode) {
			tenant = t
			return tenant, nil
		}
	}
	return tenant, errors.New("租户编码不存在")
}
