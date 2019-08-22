package events

import (
	"context"
	"encoding/json"

	"github.com/LeoAdamek/ksuid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// KinesisSink records events to an AWS Kinesis Data Stream
type KinesisSink struct {
	StreamName string
	client     *kinesis.Kinesis
}

// NewKinesisSink creates a new Kinesis event sink
func NewKinesisSink(cfg *aws.Config) (*KinesisSink, error) {

	s, err := session.NewSession(cfg)

	if err != nil {
		return nil, err
	}

	kc := kinesis.New(s)

	return &KinesisSink{
		client: kc,
	}, nil

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

	if _, err := k.client.PutRecordsWithContext(ctx, input); err != nil {
		return err
	}

	return nil
}
