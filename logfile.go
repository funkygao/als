package als

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	logfileRegex     = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)\.(\d+).\.log`)}
	dateLogfileRegex = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)_(\d+)_(.+)`)}
)

// in-flight: mongo_slow.0.log
// history: mongo_slow_20140103060105_0
type AlsLogfile struct {
	path        string // absolute path for a single file
	endsWithLog bool
}

func NewAlsLogfile() (this *AlsLogfile) {
	this = new(AlsLogfile)
	this.endsWithLog = true
	return
}

func (this *AlsLogfile) SetPath(path string) {
	this.path = path
}

func (this *AlsLogfile) SetDatePath(path string) {
	this.SetPath(path)
	this.endsWithLog = false
}

func (this *AlsLogfile) Base() string {
	return filepath.Base(this.path)
}

func (this *AlsLogfile) CamalCaseName() string {
	var m map[string]string
	if this.endsWithLog {
		m = logfileRegex.FindStringSubmatchMap(this.Base())
	} else {
		m = dateLogfileRegex.FindStringSubmatchMap(this.Base())
	}

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
