package tobject

import "time"

//Type Unit is redirection of time.Duration.
//Unit can be implemented in implementation of gadget.Gadget to control the tick rate.
type Unit time.Duration

const (
	Ms   = Unit(time.Millisecond) //Ms is Unit that refers to millisecond.
	Sec  = Unit(time.Second)      //Sec is Unit that refers to second.
	Min  = Unit(time.Minute)      //Min is Unit that refers to minute.
	Hour = Unit(time.Hour)        //Hour is Unit that refers to hour.
)
