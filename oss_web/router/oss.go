package router

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/oss_web/api"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("/presigned", api.PresignedPut)
		OssRouter.POST("/callback", api.HandlerRequest)
	}
}
