package router

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/application"
	"lzq-admin/middleware"
)

// AdminRouter 系统路由
func AdminRouter(router *gin.RouterGroup) {
	router.Use()
	{
		testRouter := router.Group("/test").Use()
		{
			testRouter.GET("/testReflectSetValue", application.ITestAppService.TestReflectSetValue)
		}

		authRouter := router.Group("/auth").Use()
		{
			authRouter.POST("/login", application.IAuthAppService.Login)
			authRouter.Use(middleware.CheckJwtToken()).POST("/logOut", application.IAuthAppService.Logout)
		}

		tenantRouter := router.Group("/tenant").Use(middleware.CheckJwtToken())
		{
			tenantRouter.POST("/create", application.ITenantAppService.Create)
		}

		systemUserRouter := router.Group("/sysUser").Use(middleware.CheckJwtToken())
		{
			systemUserRouter.POST("/sysUser", application.ISystemUserAppService.Create)
			systemUserRouter.GET("/get", application.ISystemUserAppService.Get)
			systemUserRouter.DELETE("/user", application.ISystemUserAppService.Delete)
			systemUserRouter.GET("/sysUserList", application.ISystemUserAppService.GetList)
			systemUserRouter.POST("/editSysUser", application.ISystemUserAppService.Update)
			systemUserRouter.GET("/userInfo", application.ISystemUserAppService.GetUserInfo)
			systemUserRouter.PUT("/sysUserStatus", application.ISystemUserAppService.UpdateSystemStatus)
			systemUserRouter.POST("/editUserPassword", application.ISystemUserAppService.UpdateSystemUserPassword)
			systemUserRouter.GET("/defaultAvatar", application.ISystemUserAppService.GetDefaultAvatar)
		}

		authModuleRouter := router.Group("/authModule").Use(middleware.CheckJwtToken())
		{
			authModuleRouter.POST("/create", application.IAuthModuleAppService.Create)
			authModuleRouter.GET("/get", application.IAuthModuleAppService.Get)
			authModuleRouter.DELETE("/delete", application.IAuthModuleAppService.Delete)
			authModuleRouter.GET("/list", application.IAuthModuleAppService.GetList)
			authModuleRouter.PUT("/update", application.IAuthModuleAppService.Update)
		}

		authMenuRouter := router.Group("/menu").Use(middleware.CheckJwtToken())
		{
			authMenuRouter.POST("/menu", application.IAuthMenuAppService.Create)
			authMenuRouter.GET("/get", application.IAuthMenuAppService.Get)
			authMenuRouter.DELETE("/menu", application.IAuthMenuAppService.Delete)
			authMenuRouter.GET("/menusList", application.IAuthMenuAppService.GetList)
			authMenuRouter.PUT("/menu", application.IAuthMenuAppService.Update)
			authMenuRouter.GET("/menuList", application.IAuthMenuAppService.GetMenuList)
		}

		authPermissionRouter := router.Group("/permission").Use(middleware.CheckJwtToken())
		{
			authPermissionRouter.POST("/permission", application.IAuthPermissionAppService.Create)
			authPermissionRouter.GET("/get", application.IAuthPermissionAppService.Get)
			authPermissionRouter.DELETE("/permission", application.IAuthPermissionAppService.Delete)
			authPermissionRouter.GET("/data", application.IAuthPermissionAppService.GetList)
			authPermissionRouter.PUT("/permission", application.IAuthPermissionAppService.Update)
			authPermissionRouter.GET("/permissionGroup", application.IAuthPermissionAppService.GetPermissionGroup)
		}

		authRoleRouter := router.Group("/role").Use(middleware.CheckJwtToken())
		{
			authRoleRouter.POST("/role", application.IAuthRoleAppService.Create)
			authRoleRouter.GET("/get", application.IAuthRoleAppService.Get)
			authRoleRouter.DELETE("/role", application.IAuthRoleAppService.Delete)
			authRoleRouter.GET("/roleList", application.IAuthRoleAppService.GetList)
			authRoleRouter.PUT("/role", application.IAuthRoleAppService.Update)
			authRoleRouter.PUT("/roleStatus", application.IAuthRoleAppService.UpdateRoleStatus)
			authRoleRouter.GET("/roles", application.IAuthRoleAppService.GetEnanleRoles)
		}

		authCheckerRouter := router.Group("/authenticateChecker").Use(middleware.CheckJwtToken())
		{
			authCheckerRouter.GET("/grantedMenus", application.IAuthCheckerAppService.GetGrantedMenus)
		}

		authorizeRouter := router.Group("/authorize").Use(middleware.CheckJwtToken())
		{
			authorizeRouter.GET("/rolePermissionDatas/:roleId", application.IAuthRoleAppService.GetRolePermissionDatas)
			authorizeRouter.POST("/grantPermissions", application.IAuthRoleAppService.GrantPermissions)
			authorizeRouter.DELETE("/userRole", application.IAuthCheckerAppService.DeleteUserRole)
		}
		permissionCheckerRouter := router.Group("/permissionChecker").Use(middleware.CheckJwtToken())
		{
			permissionCheckerRouter.GET("/grantedPermissions", application.IAuthCheckerAppService.GetCurrentUserGrantedPermissions)
		}
	}
}
