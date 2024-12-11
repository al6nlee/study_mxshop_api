package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"study_mxshop_api/user_web/global"
	"study_mxshop_api/user_web/initialize"
	validator2 "study_mxshop_api/user_web/validator"
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

	// 5. 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", validator2.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.PORT)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.PORT)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
