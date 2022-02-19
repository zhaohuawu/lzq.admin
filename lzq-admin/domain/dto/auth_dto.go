package dto

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/05
 * @Version 1.0.0
 */

type LoginDto struct {
	TenantCode string `json:"tenantCode" form:"tenantCode" binding:"required"` //租户编码
	LoginName  string `json:"loginName" form:"loginName" binding:"required"`   //登录名
	Password   string `json:"password" form:"password" binding:"required"`     //密码
}

type LoginTokenResponseDto struct {
	AccessToken      string      `json:"accessToken"`
	IdentityToken    string      `json:"identityToken"`
	TokenType        string      `json:"tokenType"`
	RefreshToken     string      `json:"refreshToken"`
	ErrorDescription string      `json:"errorDescription"`
	ExpiresIn        int         `json:"expiresIn"`
	Exception        interface{} `json:"exception"`
	IsError          bool        `json:"isError"`
	Name             string      `json:"name"`
	ID               string      `json:"id"`
	IsTenantAdmin    bool        `json:"isTenantAdmin"`
}
