package middleware

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/config"
	"lzq-admin/config/appsettings"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/pkg/tenant"
	"net/http"
	"strings"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/3/11
 * @Version 1.0.0
 */

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登陆系统"})
			c.Abort()
			return
		} else {
			if strings.HasPrefix(accessToken, "Bearer ") {
				if time.Now().Unix() > TokenClaims.ExpiresAt {
					c.JSON(http.StatusOK, gin.H{"msg": "Token已超时"})
					c.Abort()
					return
				}
				// 校验租户
				if config.Config.UseMultiTenancy {
					var tenantInfo appsettings.TenantInfo
					var err error
					tenantInfo, err = tenant.GetTenantById(TokenClaims.TenantId)
					if err != nil {
						c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
						c.Abort()
						return
					}
					if tenantInfo.Status != domainconsts.TenantStatusEnable {
						c.JSON(http.StatusOK, gin.H{"msg": "租户已" + domainconsts.GetConstFlag(tenantInfo.Status, domainconsts.TenantConstFlags)})
						c.Abort()
						return
					}
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "Token开通必须以Bearer+空格开头"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
