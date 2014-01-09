package als

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestBaseName(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "session_20131208230103_1", logfile.Base())
	logfile.SetPath("session_20131208230103_1")
	assert.Equal(t, "session_20131208230103_1", logfile.Base())
	logfile.SetPath("var/a.log")
	assert.Equal(t, "a.log", logfile.Base())
}

func TestLogfileCamalCaseName(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/var/bi_first_payment.10.log")
	assert.Equal(t, "biFirstPayment", logfile.CamelCaseName())

	logfile.SetPath("/var/ffs.client.Error.11.log")
	assert.Equal(t, "ffsClientError", logfile.CamelCaseName())

	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "session", logfile.CamelCaseName())

	logfile.SetPath("/data2/als/click/check_click_20140101050105_1")
	assert.Equal(t, "checkClick", logfile.CamelCaseName())

	logfile.SetPath("pv.1.log")
	assert.Equal(t, "pv", logfile.CamelCaseName())

	logfile.SetPath("/var/a/a.4.log")
	assert.Equal(t, "a", logfile.CamelCaseName())
}

func TestMatchPrefix(t *testing.T) {
	l := NewAlsLogfile()
	l.SetPath("/var////asdfasfa.log")
	assert.Equal(t, true, l.MatchPrefix("asd"))
}

func BenchmarkLogfileExt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewAlsLogfile()
		l.SetPath("/var//funplus/logs/fp_rstory/history/session_mm_20131208230103_1")
		l.Ext()
		l.SetPath("/var/bi_first_payment.10.log")
		l.Ext()
	}
}

func BenchmarkCamelCaseName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewAlsLogfile()
		l.SetPath("/var//funplus/logs/fp_rstory/history/session_mm_20131208230103_1")
		l.CamelCaseName()
		l.SetPath("/var/bi_first_payment.10.log")
		l.CamelCaseName()
	}
}
