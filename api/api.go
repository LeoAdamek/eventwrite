package api

import (
	"context"
	"net/http"
	"time"

	"bitbucket.org/mr-zen/eventwrite/events"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// API represents the application API
type API struct {
	engine *gin.Engine
	log    *zap.Logger

	buffer   *events.Buffer
	flushDst events.Sink
}

// New sets up a new instance of the API
func New(logger *zap.Logger) (*API, error) {

	a := &API{log: logger.Named("api")}

	a.buffer = events.NewBuffer()

	sink, err := events.NewKinesisSink(nil)

	if err != nil {
		return nil, err
	}

	sink.StreamName = viper.GetString("stream_name")

	a.flushDst = sink

	a.engine = gin.New()

	a.engine.Use(a.auth)

	a.engine.POST("/events", a.eventsHandler)

	go a.flushPeriodically()

	return a, nil
}

// Engine gets the HTTP mux for the server
func (a *API) Engine() http.Handler {
	return a.engine
}

// Flushes out any events that havn't been recorded every so often,
// used for periods of low activity.
func (a API) flushPeriodically() {
	t := time.NewTicker(10 * time.Second)

	for {
		<-t.C

		if ne := len(a.buffer.Events()); ne > 0 {
			a.log.Debug("Flushing Events", zap.Int("count", ne))
			if err := a.buffer.Flush(context.Background(), a.flushDst); err != nil {
				a.log.Error("Failed to flush events", zap.Error(err))
			} else {
				a.log.Debug("Flushed Events")
			}
		} else {
			a.log.Debug("No events to flush. Skipping.")
		}
	}
}
