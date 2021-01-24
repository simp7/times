package variables

import "errors"

var (
	NegativeTime = errors.New("time can't be negative") //NegativeTime is called when struct that implements tobject.Time becomes negative
)
