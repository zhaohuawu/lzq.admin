package main

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lzq-admin/config"
	"lzq-admin/config/appsettings"
	_ "lzq-admin/pkg/cache"
	"lzq-admin/pkg/hsflogger"
	_ "lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/orm"
	"lzq-admin/router"
	"net/http"
)

// @title lzq-admin Project API
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.lzqit.cn
// @contact.email lzqcode@163.com

// @host localhost:30001
// @Security x-token
// @param x-token header string true "Authorization"
// 程序入口
func main() {
	hsflogger.Init()
	// 配置初始化
	gin.SetMode(config.Config.ServerConfig.RunMode)
	// 数据库连接初始化
	orm.DatabaseInit()
	// 业务配置初始化
	appsettings.Init()
	// 装载路由
	r := router.Init()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.ServerConfig.HttpPort),
		Handler: r,
	}
	_ = server.ListenAndServe()

}
