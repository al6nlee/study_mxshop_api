package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	// logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
