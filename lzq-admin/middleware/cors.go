package middleware

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/26
 * @Version 1.0.0
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Options struct {
	Origin string
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if strings.ToLower(c.Request.Method) == "options" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
