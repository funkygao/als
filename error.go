package als

import (
	"errors"
)

var (
	ErrEmptyLine        = errors.New("empty line")
	ErrTimestampInvalid = errors.New("timestamp is not int")
	ErrEmptyArea        = errors.New("empty area")
	ErrFieldNotEnough   = errors.New("not 3 fields")
	ErrEmptyJsonPayload = errors.New("empty payload json")
)
