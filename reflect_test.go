package als

import (
	"github.com/funkygao/assert"
	"testing"
)

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
