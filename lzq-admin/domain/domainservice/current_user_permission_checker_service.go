package domainservice

import "lzq-admin/middleware"

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/27
 * @Version 1.0.0
 */

type currentUserPermissionChecker struct {

}

var CurrentUserPermissionChecker = currentUserPermissionChecker{}

func (c *currentUserPermissionChecker) IsGranted(policy string) bool {
	userId := middleware.TokenClaims.Id
	if len(userId) <= 0 {
		return false
	}
	return AuthCheckerDomainService.IsUserGranted(userId, policy)
}
