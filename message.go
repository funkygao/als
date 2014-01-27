package als

import (
	"errors"
	"fmt"
	json "github.com/funkygao/go-simplejson"
	"path/filepath"
	"reflect"
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

	size        int // in byte
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
	this.size = 0
	this.Payload = ""
	this.payloadJson = nil
}

// Timestamp will be partitially lost if in ms
func (this *AlsMessage) RawLine() string {
	return fmt.Sprintf("%s,%d,%s", this.Area, this.Timestamp, this.Payload)
}

// 26636 ns/op
func (this *AlsMessage) FromLine(line string) error {
	area, timestamp, payload, err := parseAlsLine(line)
	if err != nil {
		return err
	}

	this.size = len(line)
	this.Area = area
	this.Timestamp = timestamp
	this.Payload = payload
	this.decoded = false

	_, err = this.jsonize()
	if err != nil {
		return err
	}

	return nil
}

// 1534 ns/op
func (this *AlsMessage) FromEmptyJson() {
	this.payloadJson, _ = json.NewJson([]byte(`{}`))
	this.decoded = true
	this.size = 2
}

// TODO
func (this *AlsMessage) FromBytes(bytes []byte) error {
	return nil
}

func (this *AlsMessage) IsNilJson() bool {
	if this.Payload == "" {
		return true
	}

	js, _ := this.jsonize()
	return js.IsNil()
}

// 44534 ns/op
func (this *AlsMessage) Clone() (that *AlsMessage) {
	js, _ := this.jsonize()
	if js.IsNil() {
		return nil
	}

	that = NewAlsMessage()
	that.Area = this.Area
	that.Timestamp = this.Timestamp
	that.Payload = this.Payload
	that.Priority = this.Priority
	body, _ := js.MarshalJSON()
	that.payloadJson, _ = json.NewJson(body)
	that.decoded = true
	return
}

// 2775 ns/op
func (this *AlsMessage) QuickClone() (that *AlsMessage) {
	that = NewAlsMessage()
	that.FromEmptyJson()

	that.Area = this.Area
	that.Timestamp = this.Timestamp
	that.Payload = this.Payload
	that.Priority = this.Priority

	m, err := this.Map()
	if err != nil {
		return nil
	}

	for k, v := range m {
		that.SetField(k, v)
	}

	return
}

func (this *AlsMessage) Map() (map[string]interface{}, error) {
	js, err := this.jsonize()
	if err != nil {
		return nil, err
	}

	return js.Map()
}

func (this *AlsMessage) SetSize(size int) {
	this.size = size
}

func (this *AlsMessage) Size() int {
	return this.size
}

func (this *AlsMessage) jsonize() (data *json.Json, err error) {
	if this.decoded {
		if this.payloadJson.IsNil() {
			return this.payloadJson, ErrEmptyJsonPayload
		}

		return this.payloadJson, nil
	}

	if this.Payload == "" {
		this.FromEmptyJson()
		return
	}

	data, err = json.NewJson([]byte(this.Payload))
	if err != nil {
		data = nil
		return
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
	_, err := this.jsonize()
	if err != nil {
		return
	}

	this.payloadJson.Del(name)
}

// Does not work for _log_info.ip like deep set
func (this *AlsMessage) SetField(name string, value interface{}) (err error) {
	_, err = this.jsonize()
	if err != nil {
		return
	}

	this.payloadJson.Set(name, value) // TODO DeepSet
	return nil
}

func (this *AlsMessage) AddField(name string, value interface{}) (err error) {
	_, err = this.jsonize()
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
	_, err = this.jsonize()
	if err != nil {
		return
	}

	val = this.payloadJson.DeepGet(keyName)
	return
}

func (this *AlsMessage) NormalizedValueOfKey(keyName string) (val interface{}, err error) {
	val, err = this.ValueOfKey(keyName)
	if err != nil {
		return
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	}

	return nil, ErrUnkownType
}

// Payload field value by key name and key type
func (this *AlsMessage) FieldValue(keyName string, keyType string) (val interface{}, err error) {
	_, err = this.jsonize()
	if err != nil {
		return
	}

	jval := this.payloadJson.DeepGet(keyName)
	if jval.IsNil() {
		err = ErrEmptyJsonPayload
		return
	}

	switch keyType {
	case KEY_TYPE_STRING, KEY_TYPE_IP:
		val, err = jval.String()
	case KEY_TYPE_FLOAT:
		val, err = jval.Float64()
	case KEY_TYPE_INT, KEY_TYPE_MONEY, KEY_TYPE_RANGE:
		val, err = jval.Int()
	case KEY_TYPE_BASEFILE:
		var absoluteFilename string
		absoluteFilename, err = jval.String()
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
	js, err := this.jsonize()
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
