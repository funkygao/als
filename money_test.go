package als

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestMoneyInUsd(t *testing.T) {
	assert.Equal(t, 1445, MoneyInUsdCents("CAD", 1490))
	assert.Equal(t, 12, MoneyInUsdCents("USD", 12))
}
