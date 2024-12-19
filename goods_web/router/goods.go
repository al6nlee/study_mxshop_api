package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/goods_web/api/goods"
	"study_mxshop_api/goods_web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	// UserRouter := Router.Group("user").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("list", goods.List)
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
	}
}
