package als

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	logfileRegex     = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)\.(\d+).\.log`)}
	dateLogfileRegex = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)_(\d+)_(.+)`)}
	camelNameCache   = make(map[string]string)
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

func (this *AlsLogfile) MatchPrefix(prefix string) bool {
	return strings.HasPrefix(this.Base(), prefix)
}

func (this *AlsLogfile) CamelCaseName() string {
	md5Name := this.md5Name()
	if name, present := camelNameCache[md5Name]; present {
		return name
	}

	var m map[string]string
	if this.endsWithLog {
		m = logfileRegex.FindStringSubmatchMap(this.Base())
	} else {
		m = dateLogfileRegex.FindStringSubmatchMap(this.Base())
	}

	name := CamelCase(m["bn"])
	camelNameCache[md5Name] = name
	return name
}

func (this *AlsLogfile) md5Name() string {
	m := md5.New()
	io.WriteString(m, this.path)
	return hex.EncodeToString(m.Sum(nil))
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
