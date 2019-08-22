package events

import "context"

type Buffer struct {
	events chan Event
}

// NewBuffer creates a new events buffer.
func NewBuffer() *Buffer {

	b := &Buffer{
		events: make(chan Event, 10),
	}

	return b
}

// Events gets the event writing channel
func (b Buffer) Events() chan<- Event {
	return b.events
}

// Flush will flush out the buffer to persistent storage
func (b *Buffer) Flush(ctx context.Context, dst Sink) error {

	events := make([]Event, 10)

	// Read (up to) 10 events
	for i := 0; i < 10; i++ {
		select {
		case event := <-b.events:
			events[i] = event
		default:
			break
		}
	}

	return dst.RecordEvents(ctx, events)
}
