package als

import (
	"errors"
	"github.com/funkygao/assert"
	"testing"
)

func TestErrorEquals(t *testing.T) {
	assert.Equal(t, errors.New("empty area"), ErrEmptyArea)
}
