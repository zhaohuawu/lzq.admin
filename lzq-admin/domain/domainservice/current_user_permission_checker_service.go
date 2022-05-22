package domainservice

import (
	token "lzq-admin/pkg/auth"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/27
 * @Version 1.0.0
 */

type currentUserPermissionChecker struct {

}

var CurrentUserPermissionChecker = currentUserPermissionChecker{}

func (c *currentUserPermissionChecker) IsGranted(policy string) bool {
	//fmt.Println("GlobalTokenClaims：",token.GlobalTokenClaims)
	userId := token.GlobalTokenClaims.Id
	if len(userId) <= 0 {
		return false
	}
	return AuthCheckerDomainService.IsUserGranted(userId, policy)
}
