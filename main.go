package main

import (
	"net/http"

	"bitbucket.org/mr-zen/eventwrite/api"
	"bitbucket.org/mr-zen/eventwrite/internal/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetEnvPrefix("EW")
	viper.AutomaticEnv()
}

func main() {

	logger := logging.GetLogger()
	app, err := api.New(logger)

	if err != nil {
		logger.Fatal("Unable to bootstrap application", zap.Error(err))
	}

	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())

	go func() {
		logger.Info("Serving prometheus metrics on :9200")
		if err := http.ListenAndServe(":9200", metricsMux); err != nil {
			logger.Fatal("Unable to start metrics listener", zap.Error(err))
		}
	}()

	logger.Info("Starting service", zap.String("addr", ":8080"))
	if err := http.ListenAndServe(":8080", app.Engine()); err != nil {
		logger.Fatal("Unable to listen", zap.Error(err))
	} else {
		logger.Debug("Events flushed")
	}
}
