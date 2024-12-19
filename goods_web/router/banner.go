package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/goods_web/api/banner"
	"study_mxshop_api/goods_web/middlewares"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", banner.List)                                                            // 轮播图列表页
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.Delete) // 删除轮播图
		BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.New)          // 新建轮播图
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.Update)    // 修改轮播图信息
	}
}
