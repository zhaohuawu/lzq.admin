package middleware

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/pkg/hsflogger"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/18
 * @Version 1.0.0
 */

type RequstParams map[string]interface{}

func LogRequst() gin.HandlerFunc {
	return func(c *gin.Context) {
		requstParams := make(map[string]interface{})
		requstParams["fields.Request.Method"] = c.Request.Method
		requstParams["fields.Request.URL"] = c.Request.URL
		requstParams["fields.Request.RequestURI"] = c.Request.RequestURI
		requstParams["fields.Request.Host"] = c.Request.Host
		requstParams["fields.Request.ContentLength"] = c.Request.ContentLength
		requstParams["fields.Request.Body"] = c.Request.Body
		requstParams["fields.Request.Header"] = c.Request.Header
		hsflogger.Fields = requstParams
		hsflogger.LogInformation("123")
		c.Next()
	}
}
