package als

import (
	"encoding/gob"
	"os"
	"sync"
)

type FileCheckpoint struct {
	rwmutex *sync.RWMutex

	Files    map[string]bool // capital so that gob can serialize this part
	dumpfile string
}

func NewFileCheckpoint(dumpfile string) (this *FileCheckpoint) {
	this = new(FileCheckpoint)
	this.rwmutex = new(sync.RWMutex)
	this.dumpfile = dumpfile
	this.Files = make(map[string]bool)

	return
}

func (this *FileCheckpoint) lock() {
	this.rwmutex.Lock()
}

func (this *FileCheckpoint) unlock() {
	this.rwmutex.Unlock()
}

func (this *FileCheckpoint) rlock() {
	this.rwmutex.RLock()
}

func (this *FileCheckpoint) runlock() {
	this.rwmutex.RUnlock()
}

func (this *FileCheckpoint) Put(filename string) {
	this.lock()
	defer this.unlock()

	this.Files[filename] = true
}

func (this *FileCheckpoint) Dump() {
	this.lock()
	defer this.unlock()

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

func (this *FileCheckpoint) Load() error {
	file, err := os.Open(this.dumpfile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	decoder.Decode(this)

	return nil
}

func (this *FileCheckpoint) Contains(filename string) bool {
	this.rlock()
	defer this.runlock()

	_, found := this.Files[filename]
	return found
}
