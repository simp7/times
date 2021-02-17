package ttime

import "errors"

//ErrNegativeTime is called when struct that implements times.Time becomes negative
var (
	ErrNegativeTime = errors.New("time can't be negative")
)
