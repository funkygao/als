package als

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestMd5Logfilename(t *testing.T) {
	logfile := NewAlsLogfile()
	logfile.SetPath("/mnt/funplus/logs/fp_rstory/history/session_20131208230103_1")
	assert.Equal(t, "424a1e8cb0b7cc67d3a657bcf4784b15", logfile.md5Name())
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
