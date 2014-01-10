package als

import (
	"errors"
	"fmt"
	json "github.com/funkygao/go-simplejson"
	"path/filepath"
	"strings"
	"time"
)

type AlsMessage struct {
	Area string
	// timestamp in UTC
	Timestamp uint64
	// Textual msg/json content
	Payload  string
	Priority int

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
	this.decoded = false
	this.Priority = 0
	this.Payload = ""
	this.payloadJson = nil
}

// Timestamp will be partitially lost if in ms
func (this *AlsMessage) RawLine() string {
	return fmt.Sprintf("%s,%d,%s", this.Area, this.Timestamp, this.Payload)
}

func (this *AlsMessage) FromLine(line string) error {
	area, timestamp, payload, err := parseAlsLine(line)
	if err != nil {
		return err
	}

	this.Area = area
	this.Timestamp = timestamp
	this.Payload = payload
	this.decoded = false

	var js *json.Json
	js, err = this.PayloadJson()
	if err != nil {
		return err
	}
	if js == nil {
		return ErrEmptyJsonPayload
	}

	return nil
}

func (this *AlsMessage) FromEmptyJson() {
	this.payloadJson, _ = json.NewJson([]byte(`{}`))
	this.decoded = true
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
	if err != nil {
		data = nil
	}
	this.payloadJson = data
	this.decoded = true

	return
}

func (this *AlsMessage) FieldContains(name string, substr string) bool {
	s, _ := this.FieldValue(name, KEY_TYPE_STRING)
	return strings.Contains(s.(string), substr)
}

func (this *AlsMessage) DelField(name string) {
	_, err := this.PayloadJson()
	if err != nil {
		return
	}

	this.payloadJson.Del(name)
}

func (this *AlsMessage) SetField(name string, value interface{}) (err error) {
	_, err = this.PayloadJson()
	if err != nil {
		return
	}

	this.payloadJson.Set(name, value) // TODO DeepSet
	return nil
}

func (this *AlsMessage) AddField(name string, value interface{}) (err error) {
	_, err = this.PayloadJson()
	if err != nil {
		return
	}

	m, e := this.payloadJson.Map()
	if e != nil {
		return e
	}

	if _, present := m[name]; present {
		return errors.New(name + " already exists in Message")
	}

	this.payloadJson.Set(name, value) // TODO DeepSet
	return nil
}

func (this *AlsMessage) ValueOfKey(keyName string) (val interface{}, err error) {
	_, err = this.PayloadJson()
	if err != nil {
		return
	}

	val = this.payloadJson.DeepGet(keyName)
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
	case KEY_TYPE_INT, KEY_TYPE_MONEY, KEY_TYPE_RANGE:
		val, err = this.payloadJson.DeepGet(keyName).Int()
	case KEY_TYPE_BASEFILE:
		var absoluteFilename string
		absoluteFilename, err = this.payloadJson.DeepGet(keyName).String()
		if err != nil {
			return
		}
		val = filepath.Base(absoluteFilename)
	default:
		panic("invalid key type: " + keyType)
	}

	return
}

// _loginfo.ip -> ip
func (this *AlsMessage) leafKeyName(keyName string) string {
	parts := strings.Split(keyName, ".")
	return parts[len(parts)-1]
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
