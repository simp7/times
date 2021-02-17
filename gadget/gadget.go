package gadget

import (
	"github.com/simp7/times"
)

//Ticker do function in specific period.
//Ticker is core of other gadgets - clock, stopwatch, and timer.
type Ticker interface {
	Start(func()) //Start implement function and operate Ticker. When ticker ticks, function of parameter would be called.
	Stop()        //Stop stops Ticker ticking.
}

//Clock is an interface that returns current ttime.
type Clock interface {
	times.Gadget
}

//Stopwatch is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Stopwatch interface {
	times.Gadget
}

//Timer is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Timer interface {
	times.Gadget
}
