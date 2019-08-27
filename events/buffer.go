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

	var events []Event

	// Read (up to) 10 events
	for i := 0; i < 10; i++ {
		select {
		case event := <-b.events:
			// Only include real events.
			if event.Name != "" {
				events = append(events, event)
			}
		default:
			break
		}
	}

	// Flush events in background
	go dst.RecordEvents(ctx, events)

	return nil
}
