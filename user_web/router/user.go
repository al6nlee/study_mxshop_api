package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/user_web/api"
	"study_mxshop_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	// UserRouter := Router.Group("user").Use(middlewares.JWTAuth())
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
