package als

import (
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	baseUTC = uint64(time.Now().UTC().Unix())
)

func parseAlsLine(line string) (area string, ts uint64, msg string, err error) {
	const (
		field_splitter  = ","
		field_split_num = 3
	)

	if line == "" {
		err = ErrEmptyLine
		return
	}

	fields := strings.SplitN(line, field_splitter, field_split_num)
	if len(fields) != field_split_num {
		err = ErrFieldNotEnough
		return
	}

	area = fields[0]
	if area == "" {
		err = ErrEmptyArea
		return
	}

	ts, err = strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return
	}
	if ts > 1283931748344 {
		ts /= 1000
	} else if ts < 1262275200 {
		// 2010-01-01 00:00:00 +0800 CST
		err = ErrTimestampInvalid
		return
	} else if math.Abs(float64(ts-baseUTC)) > 5184000 {
		// 60 days gap? we never analyze data 1 month ago
		// should correct it
		ts = uint64(time.Now().UTC().Unix())
	}

	msg = fields[2]
	return
}
