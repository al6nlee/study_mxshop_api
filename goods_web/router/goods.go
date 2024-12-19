package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/goods_web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	// UserRouter := Router.Group("user").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("list", goods.GetGoodsList)
	}
}
