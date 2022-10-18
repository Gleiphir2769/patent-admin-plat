package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPackageRouter)
}

// 需认证的路由代码
func registerPackageRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Package{}
	r := v1.Group("/package").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/", api.Update)
		r.DELETE("/", api.Delete)
	}
	//r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	//{
	//	r1.PUT("/role-status", api.Update2Status)
	//	//r1.PUT("/roledatascope", api.Update2DataScope)
	//}

}
