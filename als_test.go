package als

import (
	"github.com/funkygao/assert"
	"regexp"
	"testing"
)

func TestLogfileTimeStr(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "20131208230103", logfile.LogfileTimeStr())

	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_foo_20131208230103_1")
	assert.Equal(t, "20131208230103", logfile.LogfileTimeStr())
}

func TestLogfileMonth(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "12", logfile.LogfileMonth())
	assert.Equal(t, "2013", logfile.LogfileYear())
	assert.Equal(t, "201312", logfile.LogfileYearMonth())
	assert.Equal(t, "20131208", logfile.LogfileYearMonthDate())
}

func TestLogfileBizName(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/var/bi_first_payment.10.log")
	assert.Equal(t, "biFirstPayment", logfile.BizName())

	logfile.SetPath("/var/ffs.client.Error.11.log")
	assert.Equal(t, "ffsClientError", logfile.BizName())
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

func TestGetAttr(t *testing.T) {
	type foo struct {
		Name string
		Age  int
	}

	bar := new(foo)
	bar.Name = "ping"
	bar.Age = 32
	assert.Equal(t, "ping", GetAttr(bar, "Name", ""))
	assert.Equal(t, "shit", GetAttr(bar, "non-exist", "shit"))
	assert.Equal(t, 32, GetAttr(bar, "Age", -1))
}

func TestMoneyInUsd(t *testing.T) {
	usd, _ := MoneyInUsd("CAD", 1.2)
	assert.Equal(t, 1.164, usd)
}

func TestColorize(t *testing.T) {
	assert.Equal(t, "\x1b[30m\x1b[41mhello\x1b[0m",
		Colorize([]string{"FgBlack", "BgRed"}, "hello"))
}
