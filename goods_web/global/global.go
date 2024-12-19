package global

import (
	ut "github.com/go-playground/universal-translator"
	"study_mxshop_api/goods_web/config"
	"study_mxshop_api/goods_web/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient
	NacosConfig    config.NacosConfig
)
