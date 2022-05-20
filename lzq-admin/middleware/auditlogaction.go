package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
	"lzq-admin/domain/model/extrastruct"
	"lzq-admin/pkg/hsflogger"
	"strings"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/18
 * @Version 1.0.0
 */

func LogAuditLogAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		if strings.Contains(c.Request.URL.Path, "auth/login") {
			c.Next()
		} else {
			config, err := domainservice.SysConfigDomainService.GetSysConfigCacheByCode(model.ExtraGlobalConfig, "SystemGlobalConfig")
			if err != nil {
				hsflogger.LogError("获取系统全局配置异常："+err.Error(), err)
			} else {
				var globalConfig extrastruct.SystemGlobalConfig
				if err := jsoniter.UnmarshalFromString(config, &globalConfig); err != nil {
					hsflogger.LogError("解析系统全局配置异常："+err.Error(), err)
				} else {
					if globalConfig.UseAuditLogAction {
						var logAuditLogAction model.CreateLogAuditLogActionDto
						if len(c.Request.URL.RawQuery) > 0 {
							logAuditLogAction.Parameters = c.Request.URL.RawQuery
						} else {
							data, _ := c.GetRawData()
							if len(data) > 0 {
								c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 解决gin只能读取一次body的问题
								logAuditLogAction.Parameters = string(data)
							}
						}
						defer func() {
							logAuditLogAction.HTTPMethod = c.Request.Method
							logAuditLogAction.URL = c.Request.URL.Path
							logAuditLogAction.BrowserInfo = c.Request.Header.Get("User-Agent")
							logAuditLogAction.ExecutionTime = start
							logAuditLogAction.ExecutionDuration = time.Since(start).Milliseconds()
							logAuditLogAction.HTTPStatusCode = c.Writer.Status()
							logAuditLogAction.ActionType = "Other"
							logAuditLogAction.FromSource = "lzq-admin"
							logAuditLogAction.ClientIPAddress = c.ClientIP()
							if len(c.Request.Header.Get("X-Forward-For")) > 0 {
								logAuditLogAction.ClientIPAddress = c.Request.Header.Get("X-Forward-For")
							}
							if err := recover(); err != nil {
								logAuditLogAction.Exceptions, _ = jsoniter.MarshalToString(err)
							}
							requestParams := make(map[string]interface{})
							requestParams["fields.Request.URL"] = c.Request.URL
							requestParams["fields.Request.Host"] = c.Request.Host
							requestParams["fields.Request.ContentLength"] = c.Request.ContentLength
							requestParams["fields.Request.Header"] = c.Request.Header
							logAuditLogAction.ExtraProperties = requestParams
							_ = domainservice.ILogAuditLogActionService.Insert(logAuditLogAction)
						}()
					}
				}
			}
			c.Next()
		}
	}
}
