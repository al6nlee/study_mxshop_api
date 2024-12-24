package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"study_mxshop_api/userop_web/global"
	"study_mxshop_api/userop_web/initialize"
	"study_mxshop_api/userop_web/utils"
	"study_mxshop_api/userop_web/utils/register/consul"
	"syscall"
)

func main() {
	// 1. 初始化logger
	initialize.InitLogger()

	// 2.初始化配置文件
	initialize.InitConfig()

	// 3. 初始化routers
	Router := initialize.Routers()

	// 4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 5. 初始化srv的连接
	initialize.InitSrvConn()

	debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
	if debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.PORT = port
		}
	}

	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.PORT)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.PORT)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()
	// 服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.PORT, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
