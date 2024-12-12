package global

import (
	ut "github.com/go-playground/universal-translator"
	"study_mxshop_api/user_web/config"
	"study_mxshop_api/user_web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient
)
