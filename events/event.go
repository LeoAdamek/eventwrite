package events

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeoAdamek/ksuid"
)

type ID ksuid.KSUID

func init() {

	if prt, err := ksuid.MacPartitioner(); err == nil {
		ksuid.SetPartitioner(prt)
	} else {
		fmt.Println(err)
	}

}

// Event reperesents a single event
type Event struct {
	ID ID `json:"id"`

	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`

	URL       string `json:"url"`
	Source    string `json:"source_id"`
	SessionID string `json:"session_id"`

	Metrics    map[string]float64 `json:"metrics"`
	Attributes map[string]string  `json:"attributes"`
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(ksuid.KSUID(i).String())), nil
}
