package als

import (
	json "github.com/bitly/go-simplejson"
	"time"
)

type AlsMessage struct {
	Area      string
	Timestamp uint64 // timestamp in UTC
	Payload   string // raw msg content, should be json string
	Priority  int    // set by app
}

// Convert text line to AlsMessage
func NewAlsMessage() *AlsMessage {
	this := new(AlsMessage)
	this.Reset()
	return this
}

func (this *AlsMessage) Reset() {
	this.Area = ""
	this.Timestamp = 0
	this.Priority = 0
	this.Payload = ""
}

func (this *AlsMessage) ParseLine(line string) error {
	area, timestamp, payload, err := parseAlsLine(line)
	if err != nil {
		return err
	}

	this.Area = area
	this.Timestamp = timestamp
	this.Payload = payload

	return nil
}

func (this *AlsMessage) PayloadJson() (data *json.Json, err error) {
	data, err = json.NewJson([]byte(this.Payload))

	return
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
	return time.Unix(int64(this.Timestamp), 0)
}

func (this *AlsMessage) Year() (year int) {
	year, _, _ = this.Time().Date()
	return
}

func (this *AlsMessage) Month() (month int) {
	_, m, _ := this.Time().Date()
	return int(m)
}

func (this *AlsMessage) Day() (day int) {
	_, _, day = this.Time().Date()
	return day
}
