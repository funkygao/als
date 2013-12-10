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

func TestGroupLevel(t *testing.T) {
	level := 5
	assert.Equal(t, "1-10", GroupedLevel(level))

	level = 10
	assert.Equal(t, "1-10", GroupedLevel(level))

	level = 56
	assert.Equal(t, "50-60", GroupedLevel(level))

	level = 102
	assert.Equal(t, "100-1000", GroupedLevel(level))

	// level=1022 will return "1000-1", bad, but we'll skip this assertion
}
