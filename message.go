package als

import (
	"fmt"
	json "github.com/funkygao/go-simplejson"
	"path/filepath"
	"time"
)

type AlsMessage struct {
	Area string
	// timestamp in UTC
	Timestamp uint64
	// Textual msg/json content
	Payload  string
	Priority int8

	decoded     bool
	payloadJson *json.Json
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
	this.Priority = int8(0)
	this.decoded = false
	this.Payload = ""
	this.payloadJson = nil
}

func (this *AlsMessage) FromLine(line string) error {
	area, timestamp, payload, err := parseAlsLine(line)
	if err != nil {
		return err
	}

	this.Area = area
	this.Timestamp = timestamp
	this.Payload = payload

	return nil
}

// Timestamp will be partitially lost if in ms
func (this *AlsMessage) RawLine() string {
	return fmt.Sprintf("%s,%d,%s", this.Area, this.Timestamp, this.Payload)
}

// TODO
func (this *AlsMessage) FromBytes(bytes []byte) error {
	return nil
}

func (this *AlsMessage) PayloadJson() (data *json.Json, err error) {
	if this.decoded {
		return this.payloadJson, nil
	}

	data, err = json.NewJson([]byte(this.Payload))
	this.payloadJson = data
	this.decoded = true

	return
}

// Payload field value by key name and key type
func (this *AlsMessage) FieldValue(keyName string, keyType string) (val interface{}, err error) {
	_, err = this.PayloadJson()
	if err != nil {
		return
	}

	switch keyType {
	case KEY_TYPE_STRING, KEY_TYPE_IP:
		val, err = this.payloadJson.DeepGet(keyName).String()
	case KEY_TYPE_FLOAT:
		val, err = this.payloadJson.DeepGet(keyName).Float64()
	case KEY_TYPE_INT, KEY_TYPE_MONEY, KEY_TYPE_LEVEL:
		val, err = this.payloadJson.DeepGet(keyName).Int()
	case KEY_TYPE_BASEFILE:
		var fullFilename string
		fullFilename, err = this.payloadJson.DeepGet(keyName).String()
		if err != nil {
			return
		}
		val = filepath.Base(fullFilename)
	default:
		panic("invalid key type: " + keyType)
	}

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
