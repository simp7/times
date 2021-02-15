package gadget

import (
	"github.com/simp7/times"
)

//Clock is an interface that returns current time.
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
