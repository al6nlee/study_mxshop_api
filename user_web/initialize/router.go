package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_mxshop_api/user_web/middlewares"
	"study_mxshop_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors()) // 解决跨域问题

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	ApiGroup := Router.Group("/api/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
