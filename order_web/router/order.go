package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/order_web/api"
	"study_mxshop_api/order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("", api.List)       // 订单列表
		OrderRouter.POST("", api.New)       // 新建订单
		OrderRouter.GET("/:id", api.Detail) // 订单详情
	}
}
