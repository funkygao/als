package als

import (
	"errors"
)

var (
	ErrEmptyLine        = errors.New("empty line")
	ErrTimestampInvalid = errors.New("timestamp is not int")
	ErrEmptyArea        = errors.New("empty area")
	ErrFieldNotEnough   = errors.New("not enough fields(should be 3)")
)
