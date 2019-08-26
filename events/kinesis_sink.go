package events

import (
	"context"
	"encoding/json"

	"bitbucket.org/mr-zen/eventwrite/internal/metrics"
	"github.com/LeoAdamek/ksuid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/prometheus/client_golang/prometheus"
)

// KinesisSink records events to an AWS Kinesis Data Stream
type KinesisSink struct {
	StreamName string
	client     *kinesis.Kinesis

	metrics struct {
		eventsPushedTotal prometheus.Counter
		eventsFailedTotal prometheus.Counter
	}
}

// NewKinesisSink creates a new Kinesis event sink
func NewKinesisSink(cfg *aws.Config) (*KinesisSink, error) {

	s, err := session.NewSession(cfg)

	if err != nil {
		return nil, err
	}

	kc := kinesis.New(s)

	sink := &KinesisSink{
		client: kc,
	}

	sink.metrics.eventsPushedTotal = metrics.EventsPushedTotal.WithLabelValues("kinesis")
	sink.metrics.eventsFailedTotal = metrics.EventsErrorTotal.WithLabelValues("kinesis")

	return sink, nil
}

// RecordEvents records the events to the Kinesis Stream
func (k KinesisSink) RecordEvents(ctx context.Context, events []Event) error {

	input := &kinesis.PutRecordsInput{}

	input.SetStreamName(k.StreamName)

	nEvents := len(events)
	entries := make([]*kinesis.PutRecordsRequestEntry, nEvents)

	for i, event := range events {

		data, _ := json.Marshal(event)

		entries[i] = &kinesis.PutRecordsRequestEntry{
			Data:         data,
			PartitionKey: aws.String(ksuid.KSUID(event.ID).String()),
		}
	}

	input.Records = entries

	if res, err := k.client.PutRecordsWithContext(ctx, input); err != nil {
		k.metrics.eventsFailedTotal.Add(float64(nEvents))
		return err
	} else {
		k.metrics.eventsFailedTotal.Add(float64(*res.FailedRecordCount))
		k.metrics.eventsPushedTotal.Add(float64(len(res.Records)))
	}

	return nil
}
