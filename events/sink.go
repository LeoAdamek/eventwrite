package events

import (
	"context"

	"go.uber.org/zap"
)

// Sink represents an event sink
// A sink can record events in some way
type Sink interface {
	RecordEvents(context.Context, []Event) error
}

// VoidSink discards the events.
type VoidSink struct{}

// RecordEvents discards the events
func (v VoidSink) RecordEvents(_ context.Context, _ []Event) error {
	return nil
}

// zapSink writes the events to a zapLogger
type zapSink struct {
	l *zap.Logger
}

// ZapSink creates a new zapSink
func ZapSink(log *zap.Logger) *zapSink {
	return &zapSink{l: log}
}

func (z zapSink) RecordEvents(_ context.Context, events []Event) error {

	for _, e := range events {
		z.l.Info("Event", zap.Any("event", e))
	}

	return nil
}
