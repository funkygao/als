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
