package als

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestIntsGroupLabel(t *testing.T) {
	var n = []int{1, 4, 9, 100}
	assert.Equal(t, []string{"1-4", "4-9", "9-100"}, GroupIntLabels(n))
}

func TestGroupInt(t *testing.T) {
	var ranges = []int{1, 10, 30}
	assert.Equal(t, "1-10", GroupInt(1, ranges))
	assert.Equal(t, "1-10", GroupInt(2, ranges))
	assert.Equal(t, "10-30", GroupInt(10, ranges))
	assert.Equal(t, "", GroupInt(40, ranges))
	assert.Equal(t, "", GroupInt(0, ranges))
}
