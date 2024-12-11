package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"study_mxshop_api/user_web/global"
	"study_mxshop_api/user_web/initialize"
	"syscall"
)

func main() {
	// 1. 初始化logger
	initialize.InitLogger()

	// 2. 初始化routers
	Router := initialize.Routers()

	zap.S().Debugf("启动服务器, 端口： %d", global.PORT)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.PORT)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
