package tobject

import "time"

//Unit is a type that redirection of time.Duration.
//Unit can be implemented in implementation of gadget.Gadget to control the tick rate.
type Unit time.Duration

//Ms is Unit that refers to millisecond.
//Sec is Unit that refers to second.
//Min is Unit that refers to minute.
//Hour is Unit that refers to hour.
const (
	Ms   = Unit(time.Millisecond)
	Sec  = Unit(time.Second)
	Min  = Unit(time.Minute)
	Hour = Unit(time.Hour)
)
