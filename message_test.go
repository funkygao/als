package als

import (
	"github.com/funkygao/assert"
	"github.com/funkygao/pretty"
	"testing"
)

func prepareMsgForTest() *AlsMessage {
	line := `us,1387789257065,{"uri":"\/?fb_source=canvas_bookmark","_log_info":{"uid":8664420,"snsid":"100005855171890","level":29,"gender":"female","ab":"a","payment_cash":197,"script_id":2324196651,"serial":1,"uri":"\/","host":"172.31.1.244","ip":"209.202.60.244","callee":"POST+\/+24c55bb0","sid":null, "logfile": "/var/log/a.log"}}`
	msg := NewAlsMessage()
	msg.FromLine(line)
	return msg
}

func TestAlsMessageBasic(t *testing.T) {
	msg := prepareMsgForTest()
	assert.Equal(t, "us", msg.Area)
	assert.Equal(t, uint64(1387789257065/1000), msg.Timestamp)
}

func TestAlsMessageFieldValue(t *testing.T) {
	msg := prepareMsgForTest()
	ip, err := msg.FieldValue("_log_info.ip", KEY_TYPE_IP)
	if false {
		pretty.Printf("%# v\n", *msg.payloadJson)
	}

	assert.Equal(t, nil, err)
	assert.Equal(t, "209.202.60.244", ip.(string))

	logfile, _ := msg.FieldValue("_log_info.logfile", KEY_TYPE_BASEFILE)
	assert.Equal(t, "a.log", logfile.(string))
}

func TestLeafKeyName(t *testing.T) {
	msg := prepareMsgForTest()
	assert.Equal(t, "person", msg.leafKeyName("person"))
	assert.Equal(t, "age", msg.leafKeyName("person.age"))
	assert.Equal(t, "ip", msg.leafKeyName("_log_info.ip"))
}

func TestAlsMessageTime(t *testing.T) {
	msg := prepareMsgForTest()
	year, month, day := msg.Time().Date()
	assert.Equal(t, 2013, year)
	assert.Equal(t, "December", month.String())
	assert.Equal(t, 12, int(month))
	assert.Equal(t, 23, day)
	assert.Equal(t, 23, msg.Day())
	assert.Equal(t, 12, msg.Month())
}

func TestAlsMessageJson(t *testing.T) {
	msg := prepareMsgForTest()
	json, err := msg.PayloadJson()
	assert.Equal(t, nil, err)
	uri, _ := json.Get("uri").String()
	assert.Equal(t, "/?fb_source=canvas_bookmark", uri)

	loginfo := json.Get("_log_info")
	ip, _ := loginfo.Get("ip").String()
	assert.Equal(t, "209.202.60.244", ip)
}

func TestFromEmptyJson(t *testing.T) {
	msg := NewAlsMessage()
	msg.FromEmptyJson()
	t.Logf("%#v\n", *msg.payloadJson)
	msg.SetField("foo", "bar")
	t.Logf("%#v\n", *msg.payloadJson)
	val, _ := msg.FieldValue("foo", KEY_TYPE_STRING)
	assert.Equal(t, "bar", val)
}
