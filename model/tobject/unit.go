package tobject

import "time"

type Unit time.Duration

const (
	Ms = Unit(time.Millisecond)
	Sec = Unit(time.Second)
	Min = Unit(time.Minute)
	Hour = Unit(time.Hour)
)