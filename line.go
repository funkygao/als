package als

import (
	"errors"
	"strconv"
	"strings"
)

func parseAlsLine(line string) (area string, ts uint64, msg string, err error) {
	const (
		field_splitter  = ","
		field_split_num = 3
	)

	fields := strings.SplitN(line, field_splitter, field_split_num)
	if len(fields) != field_split_num {
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
