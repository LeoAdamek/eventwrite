package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GetLogger gets a logger
func GetLogger() *zap.Logger {
	var logger *zap.Logger
	if viper.GetString("log_mode") == "development" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	return logger
}
