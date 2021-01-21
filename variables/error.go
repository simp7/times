package variables

import "errors"

var (
	NegativeTime = errors.New("time can't be negative")
)
