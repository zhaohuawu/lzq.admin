package router

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "lzq-admin/docs"
	"lzq-admin/middleware"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.LogAuditLogAction())

	adminRouter := router.Group("/api/app")

	/***api路由定义****/
	AdminRouter(adminRouter)
	return router
}
