package als

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type alsReader struct {
	filename string // absolute path for a single file
	stream   *os.File
	reader   *bufio.Reader
}

func NewAlsReader(filename string) (this *alsReader) {
	this = new(alsReader)
	this.filename = filename

	return
}

func (this *alsReader) Close() (err error) {
	if this.stream == nil {
		err = errors.New("must call Start before Close")
		return
	}

	this.stream.Close()
	return
}

func (this *alsReader) Start() (err error) {
	this.stream, err = os.Open(this.filename)
	if err != nil {
		return
	}

	this.reader = bufio.NewReader(this.stream)
	return
}

func (this *alsReader) ReadLine() ([]byte, error) {
	line, isPrefix, err := this.reader.ReadLine()
	if !isPrefix {
		return line, err
	}

	buf := append([]byte(nil), line...)
	for isPrefix && err == nil {
		line, isPrefix, err = this.reader.ReadLine()
		buf = append(buf, line...)
	}

	return buf, err
}

// Get time info from filename
func (this *alsReader) LogfileTimeStr() string {
	fields := strings.Split(filepath.Base(this.filename), "_")
	return fields[len(fields)-2]
}

func (this *alsReader) LogfileMonth() string {
	ts := this.LogfileTimeStr()
	return ts[4:6]
}

func (this *alsReader) LogfileYear() string {
	ts := this.LogfileTimeStr()
	return ts[:4]
}

func (this *alsReader) LogfileYearMonth() string {
	ts := this.LogfileTimeStr()
	return ts[:6]
}

func (this *alsReader) LogfileYearMonthDate() string {
	ts := this.LogfileTimeStr()
	return ts[:8]
}
