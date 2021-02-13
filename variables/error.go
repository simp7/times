package variables

import "errors"

//ErrNegativeTime is called when struct that implements tobject.Time becomes negative
var (
	ErrNegativeTime = errors.New("time can't be negative")
)
