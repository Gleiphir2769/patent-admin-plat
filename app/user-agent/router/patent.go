package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatentRouter)
}

// 需认证的路由代码
func registerPatentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Patent{}

	r := v1.Group("/patent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		//r.GET("", api.GetPatentLists)                //显示专利本地数据√
		r.GET("/:patent_id", api.GetPatentById)      //查询专利√
		r.GET("/claim", api.GetClaimPages)           //显示认领专利√
		r.GET("/focus", api.GetFocusPages)           //显示关注专利√
		r.GET("/tag-patents/:tag_id", api.GetPatent) //显示该标签下的专利√
		r.GET("/tags/:patent_id", api.GetTags)       //显示专利的标签√
		r.GET("/user", api.GetUserPatentsPages)      //获取该用户所有专利列表
		//r.GET("/relationship", api.GetTheGraphByUserId)   //获取用户关注的专利的关系图(点位置固定但点的大小数量和线的数量、值根据数据生成)
		//r.GET("/relationship2", api.GetTheGraphByUserId2) //获取用户关注的专利的关系图(点位置随机、大小数量和线的数量、值根据数据生成)

		r.POST("", api.InsertIfAbsent)    //添加专利√
		r.POST("/claim", api.ClaimPatent) //认领专利√
		r.POST("/focus", api.FocusPatent) //关注专利√
		r.POST("/tag", api.InsertTag)     //为专利添加标签√

		r.PUT("", api.UpdatePatent) //修改专利

		r.DELETE("/:patent_id", api.DeletePatent)                  //删除该专利√
		r.DELETE("/claim/:patent_id", api.DeleteClaim)             //取消关注√
		r.DELETE("/focus/:patent_id", api.DeleteFocus)             //取消认领√
		r.DELETE("/tags/:tag_id/patent/:patent_id", api.DeleteTag) //取消添加该标签√

	}

}
