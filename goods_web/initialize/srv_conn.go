package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"study_mxshop_api/goods_web/global"
	"study_mxshop_api/goods_web/proto"
)

func InitSrvConn() {
	userConn, err := grpc.Dial(
		fmt.Sprintf(
			"consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.GoodsSrv.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 轮询
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	// 生成grpc的client调用接口
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
