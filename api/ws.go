package api

import (
	"encoding/json"
	"fmt"
	"io"

	"bitbucket.org/mr-zen/eventwrite/events"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

// MessageType is an enum representing a type of message
type MessageType string

const (
	// MessageTypeEvent is the message type for a single event
	MessageTypeEvent MessageType = "event"
)

// All messages on the socket are wrapped with a type so we know
// how to unmarhsall them.
type socketMessage struct {
	MessageType MessageType     `json:"type"`
	Value       json.RawMessage `json:"msg"`
}

// socketHandler handles events streaming in by websocket
func (a API) socketHandler(w *websocket.Conn) {

	dec := json.NewDecoder(w)

	// Decode JSON objects as they come in.
	for {
		m := &socketMessage{}
		err := dec.Decode(m)

		if err != nil {
			if err == io.EOF {
				a.log.Info("Closing connection")
				w.Close()
				break
			}

			a.log.Error("Unable to read incoming message", zap.Error(err))
		}

		if err := handleMessage(m); err != nil {
			a.log.Error("Unable to process message", zap.Error(err), zap.String("type", string(m.MessageType)))
		}
	}
}

func handleMessage(m *socketMessage) error {

	switch m.MessageType {
	case MessageTypeEvent:
		ev := &events.Event{}

		err := json.Unmarshal(m.Value, ev)

		if err != nil {
			return err
		}

		break
	}

	return errUnknownMessageType(m.MessageType)
}

func errUnknownMessageType(typ MessageType) error {
	return fmt.Errorf("Unknown message type: %s", typ)
}
