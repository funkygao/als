package als

import (
	"github.com/bmizerany/assert"
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
}
