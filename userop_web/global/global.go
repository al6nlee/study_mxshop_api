package global

import (
	ut "github.com/go-playground/universal-translator"
	"study_mxshop_api/userop_web/config"
	"study_mxshop_api/userop_web/proto"
)

var (
	Trans        ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig  config.NacosConfig

	GoodsSrvClient proto.GoodsClient
	MessageClient  proto.MessageClient
	AddressClient  proto.AddressClient
	UserFavClient  proto.UserFavClient
)
