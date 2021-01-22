package tobject

import "time"

//Type Unit is redirection of time.Duration.
//Unit can be implemented in implementation of gadget.Gadget to control the tick rate.
type Unit time.Duration

const (
	Ms   = Unit(time.Millisecond)
	Sec  = Unit(time.Second)
	Min  = Unit(time.Minute)
	Hour = Unit(time.Hour)
)
