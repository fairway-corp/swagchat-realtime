package utils

import (
	"os"

	"go.uber.org/zap"
)

var AppLogger *zap.Logger

func setupLogger() {
	if AppLogger == nil {
		var err error
		var logger *zap.Logger
		if Cfg.ApiServer.LoggingLevel == "development" {
			logger, err = zap.NewDevelopment()
		} else if Cfg.ApiServer.LoggingLevel == "production" {
			logger, err = zap.NewProduction()
		} else {
			os.Exit(0)
		}
		if err != nil {
			os.Exit(0)
		}
		AppLogger = logger.WithOptions(zap.Fields(
			zap.String("appName", APP_NAME),
			zap.String("version", API_VERSION),
		))
	}
}
