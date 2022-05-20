package middleware

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/pkg/auth"
	"net/http"
	"strings"
)

var TokenClaims = &token.TokenClaims{}

func CheckJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if len(accessToken) > 0 {
			if strings.HasPrefix(accessToken, "Bearer ") {
				accessToken = strings.TrimPrefix(accessToken, "Bearer ")
				var err error
				TokenClaims, err = token.ParseToken(accessToken)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{"msg": "Token解析失败", "error": err.Error()})
					c.Abort()
					return
				}
				token.GlobalTokenClaims = TokenClaims
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "Token开通必须以Bearer+空格开头"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
