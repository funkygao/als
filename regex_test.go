package als

import (
	"github.com/funkygao/assert"
	"regexp"
	"testing"
)

func TestNamedRegexp(t *testing.T) {
	var myExp = NamedRegexp{regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)}
	m := myExp.FindStringSubmatchMap("1234.5678.9")
	assert.Equal(t, "1234", m["first"])
	assert.Equal(t, "9", m["second"])
}
