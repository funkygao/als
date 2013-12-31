package als

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAlsMessage(t *testing.T) {
	line := `us,1387789257065,{"uri":"\/?fb_source=canvas_bookmark","_log_info":{"uid":8664420,"snsid":"100005855171890","level":29,"gender":"female","ab":"a","payment_cash":197,"script_id":2324196651,"serial":1,"uri":"\/","host":"172.31.1.244","ip":"209.202.60.244","callee":"POST+\/+24c55bb0","sid":null}}`
	msg := NewAlsMessage()
	err := msg.ParseLine(line)
	msg.Priority = 5
	assert.Equal(t, nil, err)
	assert.Equal(t, "us", msg.Area)
	assert.Equal(t, 5, msg.Priority)
	assert.Equal(t, uint64(1387789257065/1000), msg.Ts)

	year, month, day := msg.Time().Date()
	assert.Equal(t, 2013, year)
	assert.Equal(t, "December", month.String())
	assert.Equal(t, 12, int(month))
	assert.Equal(t, 23, day)

	json, err := msg.PayloadJson()
	assert.Equal(t, nil, err)
	uri, _ := json.Get("uri").String()
	assert.Equal(t, "/?fb_source=canvas_bookmark", uri)

	loginfo := json.Get("_log_info")
	ip, _ := loginfo.Get("ip").String()
	assert.Equal(t, "209.202.60.244", ip)
}
