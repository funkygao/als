package als

import (
	json "github.com/bitly/go-simplejson"
)

type Message struct {
	Area    string
	Ts      uint64 // timestamp in UTC
	Payload string // raw msg content, should be json string
}

// Convert text line to Message
func NewMessage(line string) (*Message, error) {
	area, ts, payload, err := ParseAlsLine(line)
	if err != nil {
		return nil, err
	}

	msg := &Message{Area: area, Ts: ts, Payload: payload}
	return msg, nil
}

func (this *Message) PayloadJson() (*json.Json, error) {
	return MsgToJson(this.Payload)
}

func (this *Message) MarshalPayload() ([]byte, error) {
	js, err := this.PayloadJson()
	if err != nil {
		return nil, err
	}

	return js.MarshalJSON()
}
