package api

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/mr-zen/eventwrite/events"
	"github.com/LeoAdamek/ksuid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type eventsRequest struct {
	Events []events.Event `json:"events"`
}

func (a API) eventsHandler(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	e := &eventsRequest{}

	if err := decoder.Decode(e); err != nil {
		a.log.Error("Unable to decode event", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, ev := range e.Events {

		ev.ID = events.ID(ksuid.Next())
		ev.Source = c.MustGet("source_id").(string)

		a.log.Debug("Got event", zap.Any("event", e))

		select {
		case a.buffer.Events() <- ev:
		default:
			a.log.Debug("Event buffer is full. Flushing first")

			if err := a.buffer.Flush(c, a.flushDst); err != nil {
				a.log.Error("Unable to flush events", zap.Error(err))
			}

			a.buffer.Events() <- ev

		}
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "OK"})
}
