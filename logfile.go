package als

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	logfileRegex     = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)\.(\d+)\.log`)}
	dateLogfileRegex = NamedRegexp{regexp.MustCompile(`(?P<bn>.+)_(\d+)_(.+)`)}

	camelNameCache = make(map[string]string)
	cacheRWMutex   = new(sync.RWMutex)
)

// in-flight: mongo_slow.0.log
// history: mongo_slow_20140103060105_0
type AlsLogfile struct {
	path     string // absolute path for a single file
	baseName string // to lessen CPU cycle
}

func NewAlsLogfile() (this *AlsLogfile) {
	this = new(AlsLogfile)
	return
}

func (this *AlsLogfile) SetPath(path string) {
	this.path = path
	this.baseName = filepath.Base(this.path)
}

func (this *AlsLogfile) Base() string {
	return this.baseName
}

func (this *AlsLogfile) MatchPrefix(prefix string) bool {
	return strings.HasPrefix(this.Base(), prefix)
}

// FIXME  the ugly coding
func (this *AlsLogfile) CamelCaseName() string {
	md5Name := this.md5Name()
	cacheRWMutex.RLock()
	defer cacheRWMutex.RUnlock()
	if name, present := camelNameCache[md5Name]; present {
		return name
	}

	ext := filepath.Ext(this.Base())
	var m map[string]string
	if ext == ".log" {
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
