package als

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestParseValidAlsLine(t *testing.T) {
	line := `us,1387789257065,{"uri":"\/?fb_source=canvas_bookmark","_log_info":{"uid":8664420,"snsid":"100005855171890","level":29,"gender":"female","ab":"a","payment_cash":197,"script_id":2324196651,"serial":1,"uri":"\/","host":"172.31.1.244","ip":"209.202.60.244","callee":"POST+\/+24c55bb0","sid":null, "logfile": "/var/log/a.log"}}`
	area, ts, _, err := parseAlsLine(line)
	assert.Equal(t, "us", area)
	assert.Equal(t, uint64(1387789257065/1000), ts)
	assert.Equal(t, nil, err)
}

func TestParseInvalidAlsLine(t *testing.T) {
	var line string
	line = `,1387789257065,foo`
	_, _, _, err := parseAlsLine(line)
	assert.Equal(t, ErrEmptyArea, err)

	line = ``
	_, _, _, err = parseAlsLine(line)
	assert.Equal(t, ErrEmptyLine, err)

	line = `us,45,bar`
	_, _, _, err = parseAlsLine(line)
	assert.Equal(t, ErrTimestampInvalid, err)

	line = `us,1387789257065`
	_, _, _, err = parseAlsLine(line)
	assert.Equal(t, ErrFieldNotEnough, err)
}
