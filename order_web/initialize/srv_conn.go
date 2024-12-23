package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"study_mxshop_api/order_web/global"
	"study_mxshop_api/order_web/proto"
)

func InitSrvConn() {
	goodsConn, err := grpc.Dial(
		fmt.Sprintf(
			"consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.GoodsSrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 轮询
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}
	// 生成grpc的client调用接口
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	// 初始化库存服务连接
	invConn, err := grpc.Dial(
		fmt.Sprintf(
			"consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.InventorySrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【库存服务失败】")
	}
	// 生成grpc的client调用接口
	global.InventorySrvClient = proto.NewInventoryClient(invConn)

	// 初始化订单服务连接
	orderConn, err := grpc.Dial(
		fmt.Sprintf(
			"consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.OrderSrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【订单服务失败】")
	}
	// 生成grpc的client调用接口
	global.OrderSrvClient = proto.NewOrderClient(orderConn)
}
