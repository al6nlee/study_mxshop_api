package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/minio/minio-go/v7"
	"study_mxshop_api/oss_web/config"
)

var (
	Trans        ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig  config.NacosConfig
	MinioClient  *minio.Client
)
