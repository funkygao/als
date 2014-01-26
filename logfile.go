package als

import (
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	logfileRegex     = regexp.MustCompile(`(.+)\.(\d+)\.log`)
	dateLogfileRegex = regexp.MustCompile(`(.+)_(\d+)_(.+)`)

	camelNameCache        = make(map[string]string)
	camelNameCacheRWMutex = new(sync.RWMutex)
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

func (this *AlsLogfile) Path() string {
	return this.path
}

func (this *AlsLogfile) SetPath(path string) {
	this.path = path
	this.baseName = filepath.Base(this.path)
}

func (this *AlsLogfile) Base() string {
	return this.baseName
}

func (this *AlsLogfile) Ext() string {
	return filepath.Ext(this.Base())
}

func (this *AlsLogfile) MatchPrefix(prefix string) bool {
	return strings.HasPrefix(this.Base(), prefix)
}

func (this *AlsLogfile) CamelCaseName() string {
	camelNameCacheRWMutex.RLock()
	if name, present := camelNameCache[this.Base()]; present {
		camelNameCacheRWMutex.RUnlock()
		return name
	}
	camelNameCacheRWMutex.RUnlock()

	var name string
	if this.Ext() == ".log" {
		name = logfileRegex.FindStringSubmatch(this.Base())[1]
	} else {
		name = dateLogfileRegex.FindStringSubmatch(this.Base())[1]
	}

	name = CamelCase(name)
	camelNameCacheRWMutex.Lock()
	camelNameCache[this.Base()] = name
	camelNameCacheRWMutex.Unlock()
	return name
}
