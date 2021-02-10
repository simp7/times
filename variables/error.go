package variables

import "errors"

var (
	ErrNegativeTime = errors.New("time can't be negative") //ErrNegativeTime is called when struct that implements tobject.Time becomes negative
)
