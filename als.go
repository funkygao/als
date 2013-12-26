package als

import (
	"errors"
	json "github.com/bitly/go-simplejson"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	LINE_SPLITTER  = ","
	LINE_SPLIT_NUM = 3
)

func ParseAlsLine(line string) (area string, ts uint64, msg string, err error) {
	fields := strings.SplitN(line, LINE_SPLITTER, LINE_SPLIT_NUM)
	if len(fields) != LINE_SPLIT_NUM {
		err = errors.New("not enough fields: " + line)
		return
	}

	area = fields[0]
	if area == "" {
		err = errors.New("empty area")
		return
	}

	ts, err = strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return
	}
	if ts > 1283931748344 {
		ts /= 1000
	} else if ts < 1262275200 {
		err = errors.New("invalid timestamp: " + fields[1])
		return
	}

	msg = fields[2]
	return
}

func MsgToJson(msg string) (data *json.Json, err error) {
	data, err = json.NewJson([]byte(msg))

	return
}

// Extract time info from als filename
func LogfileTimeStr(filename string) string {
	fields := strings.Split(filepath.Base(filename), "_")
	return fields[len(fields)-2]
}
