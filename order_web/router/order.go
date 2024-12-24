package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/order_web/api/order"
	"study_mxshop_api/order_web/api/pay"
	"study_mxshop_api/order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}

	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify) // 支付宝支付成功后的回调函数
	}
}
