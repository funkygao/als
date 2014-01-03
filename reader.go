package als

import (
	"bufio"
	"errors"
	"os"
)

type alsReader struct {
	*alsLogfile

	stream *os.File
	reader *bufio.Reader
}

func NewAlsReader(path string) (this *alsReader) {
	this = new(alsReader)
	this.alsLogfile = NewAlsLogfile()
	this.alsLogfile.SetPath(path)

	return
}

func (this *alsReader) Close() (err error) {
	if this.stream == nil {
		err = errors.New("must call Open before Close")
		return
	}

	this.stream.Close()
	return
}

func (this *alsReader) Open() (err error) {
	this.stream, err = os.Open(this.path)
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
