package als

import (
	"encoding/gob"
	"os"
	"sync"
)

type fileCheckpoint struct {
	*sync.RWMutex

	Files    map[string]bool // capital so that gob can serialize this part
	dumpfile string
}

func NewFileCheckpoint(dumpfile string) (this *fileCheckpoint) {
	this = new(fileCheckpoint)
	this.RWMutex = new(sync.RWMutex)
	this.dumpfile = dumpfile
	this.Files = make(map[string]bool)

	return
}

func (this *fileCheckpoint) Put(filename string) {
	this.Lock()
	defer this.Unlock()

	this.Files[filename] = true
}

func (this *fileCheckpoint) Dump() {
	file, err := os.OpenFile(this.dumpfile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(*this); err != nil {
		panic(err)
	}
}

func (this *fileCheckpoint) Load() error {
	file, err := os.Open(this.dumpfile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	decoder.Decode(this)

	return nil
}

func (this *fileCheckpoint) Contains(filename string) bool {
	this.RLock()
	defer this.RUnlock()

	_, found := this.Files[filename]
	return found
}
