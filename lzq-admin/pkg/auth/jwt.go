package token

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"github.com/dgrijalva/jwt-go"
	"lzq-admin/config"
	"time"
)

type TokenClaims struct {
	LoginName string `json:"loginName"`
	Name      string `json:"name"`
	SysType   string `json:"sysType"`
	TenantId  string `json:"tenantId"`
	jwt.StandardClaims
}

var jwtSecret = []byte(config.Config.JwtConfig.JwtSecret)

const (
	SysTypeAdmin = "admin"
	SysTypeWeb   = "web"
)

//GenerateToken 签发用户Token
func GenerateToken(userId, loginName, userName, sysType string, tenantId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Config.JwtConfig.JwtExpireDate*24) * time.Hour)

	claims := TokenClaims{
		loginName,
		userName,
		sysType,
		"",
		jwt.StandardClaims{
			Id:        userId,
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Config.JwtConfig.JwtIssuer,
		},
	}
	if config.Config.ServerConfig.UseMultiTenancy {
		claims.TenantId = tenantId
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := tokenClaims.SignedString(jwtSecret)

	return accessToken, err
}

// ParseToken 解析Token
func ParseToken(accessToken string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
