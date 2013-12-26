package als

import (
	"github.com/bmizerany/assert"
	"regexp"
	"testing"
)

func TestLogfileTimeStr(t *testing.T) {
	reader := NewAlsReader("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "20131208230103", reader.LogfileTimeStr())

	reader = NewAlsReader("/mnt/funplus/logs/fp_rstory/history/session_foo_20131208230103_1")
	assert.Equal(t, "20131208230103", reader.LogfileTimeStr())
}

func TestLogfileMonth(t *testing.T) {
	reader := NewAlsReader("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "12", reader.LogfileMonth())
	assert.Equal(t, "2013", reader.LogfileYear())
	assert.Equal(t, "201312", reader.LogfileYearMonth())
	assert.Equal(t, "20131208", reader.LogfileYearMonthDate())
}

func TestIntsGroupLabel(t *testing.T) {
	var n = []int{1, 4, 9, 100}
	assert.Equal(t, []string{"1-4", "4-9", "9-100"}, GroupIntLabels(n))
}

func TestNamedRegexp(t *testing.T) {
	var myExp = NamedRegexp{regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)}
	m := myExp.FindStringSubmatchMap("1234.5678.9")
	assert.Equal(t, "1234", m["first"])
	assert.Equal(t, "9", m["second"])
}

func TestGroupInt(t *testing.T) {
	var ranges = []int{1, 10, 30}
	assert.Equal(t, "1-10", GroupInt(1, ranges))
	assert.Equal(t, "1-10", GroupInt(2, ranges))
	assert.Equal(t, "10-30", GroupInt(10, ranges))
	assert.Equal(t, "", GroupInt(40, ranges))
	assert.Equal(t, "", GroupInt(0, ranges))
}

func TestCardinalityCounter(t *testing.T) {
	c := NewCardinalityCounter()
	c.Add("dau", 34343434)
	c.Add("dau", 45454)
	c.Add("dau", 888)
	assert.Equal(t, uint64(3), c.Count("dau"))

	c.Reset("msg")
	c.Add("msg", "we are in China")
	c.Add("msg", "where are you")
	assert.Equal(t, uint64(2), c.Count("msg"))
}
