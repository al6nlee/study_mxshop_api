package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/goods_web/api/goods"
	"study_mxshop_api/goods_web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("list", goods.List)                                                             // 获取商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)               // 创建商品
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)         // 更新商品
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus) // 更新商品个别属性状态
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)      // 删除商品

		GoodsRouter.GET("/:id", goods.Detail)        // 获取商品的详情
		GoodsRouter.GET("/:id/stocks", goods.Stocks) // 获取商品的库存
	}
}
