package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/oss_web/api"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("token", api.Token)
		OssRouter.POST("/callback", api.HandlerRequest)
	}
}
