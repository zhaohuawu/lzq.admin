package middleware

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
	"net/http"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/19
 * @Version 1.0.0
 */

func LoginAuditLogAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var logAuditLogAction model.CreateLogAuditLogActionDto
		defer func() {
			logAuditLogAction.HTTPMethod = c.Request.Method
			logAuditLogAction.URL = c.Request.URL.Path
			logAuditLogAction.BrowserInfo = c.Request.Header.Get("User-Agent")
			logAuditLogAction.ExecutionTime = start
			logAuditLogAction.ExecutionDuration = time.Since(start).Milliseconds()
			logAuditLogAction.HTTPStatusCode = c.Writer.Status()
			logAuditLogAction.ActionType = "Login"
			logAuditLogAction.FromSource = "lzq-admin"
			logAuditLogAction.ClientIPAddress = c.ClientIP()
			if len(c.Request.Header.Get("X-Forward-For")) > 0 {
				logAuditLogAction.ClientIPAddress = c.Request.Header.Get("X-Forward-For")
			}
			if err := recover(); err != nil {
				logAuditLogAction.Exceptions, _ = jsoniter.MarshalToString(err)
				logAuditLogAction.HTTPStatusCode = http.StatusInternalServerError
			}
			requestParams := make(map[string]interface{})
			requestParams["fields.Request.URL"] = c.Request.URL
			requestParams["fields.Request.Host"] = c.Request.Host
			requestParams["fields.Request.ContentLength"] = c.Request.ContentLength
			requestParams["fields.Request.Header"] = c.Request.Header
			logAuditLogAction.ExtraProperties = requestParams
			_ = domainservice.ILogAuditLogActionService.Insert(logAuditLogAction)
		}()

		c.Next()

	}
}
