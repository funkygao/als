package als

import (
	"path/filepath"
	"regexp"
	"strings"
)

// in-flight: mongo_slow.0.log
// history: mongo_slow_20140103060105_0
type alsLogfile struct {
	path string // absolute path for a single file
	r    NamedRegexp
}

func NewAlsLogfile() (this *alsLogfile) {
	this = new(alsLogfile)
	this.r = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)\.(\d+).\.log`)}
	return
}

func (this *alsLogfile) SetPath(path string) {
	this.path = path
}

func (this *alsLogfile) Base() string {
	return filepath.Base(this.path)
}

func (this *alsLogfile) BizName() string {
	m := this.r.FindStringSubmatchMap(this.Base())
	return CamelCase(m["bn"])
}

// Get time info from filename
func (this *alsLogfile) LogfileTimeStr() string {
	fields := strings.Split(filepath.Base(this.path), "_")
	return fields[len(fields)-2]
}

func (this *alsLogfile) LogfileMonth() string {
	ts := this.LogfileTimeStr()
	return ts[4:6]
}

func (this *alsLogfile) LogfileYear() string {
	ts := this.LogfileTimeStr()
	return ts[:4]
}

func (this *alsLogfile) LogfileYearMonth() string {
	ts := this.LogfileTimeStr()
	return ts[:6]
}

func (this *alsLogfile) LogfileYearMonthDate() string {
	ts := this.LogfileTimeStr()
	return ts[:8]
}
