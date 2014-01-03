package als

import (
	"path/filepath"
	"regexp"
	"strings"
)

// in-flight: mongo_slow.0.log
// history: mongo_slow_20140103060105_0
type AlsLogfile struct {
	path string // absolute path for a single file
	r    NamedRegexp
}

func NewAlsLogfile() (this *AlsLogfile) {
	this = new(AlsLogfile)
	this.r = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)\.(\d+).\.log`)}
	return
}

func (this *AlsLogfile) SetPath(path string) {
	this.path = path
}

func (this *AlsLogfile) Base() string {
	return filepath.Base(this.path)
}

func (this *AlsLogfile) BizName() string {
	m := this.r.FindStringSubmatchMap(this.Base())
	return CamelCase(m["bn"])
}

// Get time info from filename
func (this *AlsLogfile) LogfileTimeStr() string {
	fields := strings.Split(filepath.Base(this.path), "_")
	return fields[len(fields)-2]
}

func (this *AlsLogfile) LogfileMonth() string {
	ts := this.LogfileTimeStr()
	return ts[4:6]
}

func (this *AlsLogfile) LogfileYear() string {
	ts := this.LogfileTimeStr()
	return ts[:4]
}

func (this *AlsLogfile) LogfileYearMonth() string {
	ts := this.LogfileTimeStr()
	return ts[:6]
}

func (this *AlsLogfile) LogfileYearMonthDate() string {
	ts := this.LogfileTimeStr()
	return ts[:8]
}
