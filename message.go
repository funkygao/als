package als

import (
	json "github.com/bitly/go-simplejson"
	"time"
)

type AlsMessage struct {
	Area     string
	Ts       uint64 // timestamp in UTC
	Payload  string // raw msg content, should be json string
	Priority int    // set by app
}

// Convert text line to AlsMessage
func NewAlsMessage(line string, priority int) (*AlsMessage, error) {
	area, ts, payload, err := ParseAlsLine(line)
	if err != nil {
		return nil, err
	}

	return &AlsMessage{Area: area, Ts: ts, Payload: payload, Priority: priority}, nil
}

func (this *AlsMessage) PayloadJson() (*json.Json, error) {
	return MsgToJson(this.Payload)
}

func (this *AlsMessage) MarshalPayload() ([]byte, error) {
	js, err := this.PayloadJson()
	if err != nil {
		return nil, err
	}

	return js.MarshalJSON()
}

// Convert timestamp from uint64 to struct
func (this *AlsMessage) Time() time.Time {
	return time.Unix(int64(this.Ts), 0)
}
