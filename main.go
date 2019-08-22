package main

import (
	"net/http"

	"bitbucket.org/mr-zen/eventwrite/api"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetEnvPrefix("EW")
	viper.AutomaticEnv()
}

func main() {
	logger, _ := zap.NewDevelopment()

	app, err := api.New(logger)

	if err != nil {
		logger.Fatal("Unable to bootstrap application", zap.Error(err))
	}

	logger.Info("Starting service", zap.String("addr", ":8885"))

	if err := http.ListenAndServe(":8885", app.Engine()); err != nil {
		logger.Fatal("Unable to listen", zap.Error(err))
	} else {
		logger.Debug("Events flushed")
	}
}
