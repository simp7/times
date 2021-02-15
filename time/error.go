package time

import "errors"

//ErrNegativeTime is called when struct that implements time.Time becomes negative
var (
	ErrNegativeTime = errors.New("time can't be negative")
)
