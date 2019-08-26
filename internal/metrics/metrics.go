package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	EventsReceivedTotal prometheus.Counter
	EventsPushedTotal   *prometheus.CounterVec
	EventsErrorTotal    *prometheus.CounterVec
)

const namespace = "eventwrite"

func init() {

	EventsReceivedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "events_received_total",
		Help:      "Total Events Received",
	})

	EventsPushedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "events_pushed_total",
		Help:      "Total Events Pushed to Sink destination",
	}, []string{"sink"})

	EventsErrorTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "events_error_total",
		Help:      "Total Events Failed to Sink destination",
	}, []string{"sink"})
}
