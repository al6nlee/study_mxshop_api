package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"strconv"
	"study_mxshop_api/user_web/global"
	"study_mxshop_api/user_web/proto"
)

func InitSrvConn() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrv.Name))
	if err != nil {
		panic(err)
	}
	userSrvHost := ""
	userSrvPort := 0
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}
	userConn, err := grpc.Dial(userSrvHost+":"+strconv.Itoa(userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接srv失败", "err", err)
	}
	// 生成grpc的client调用接口
	global.UserSrvClient = proto.NewUserClient(userConn)
}
