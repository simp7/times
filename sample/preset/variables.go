package preset

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/object"
)

//GetStandard returns object.Standard object that are set to seconds by parameter sec.
//Notice: GetStandard is just for sample. We recommend you to use object.Standard function as a substitute.
func GetStandard(sec int) times.Object {
	return object.StandardZero().SetSecond(sec)
}

//GetAccurate returns object.Accurate object that are set to seconds by parameter sec.
//Notice: GetAccurate is just for sample. We recommend you to use object.Accurate function as a substitute.
func GetAccurate(sec int) times.Object {
	return object.AccurateZero().SetSecond(sec)
}

//Stopwatches returns an array of gadget.Stopwatch.
//Notice: Stopwatches is just for sample.
func Stopwatches() []times.Gadget {
	return []times.Gadget{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail}
}

//Timers returns an array of gadget.Timer that seconds are set to value of parameter deadline.
//Notice: Timers is just for sample.
func Timers(deadline int) []times.Gadget {
	return []times.Gadget{timer.Standard(GetStandard(deadline)), timer.Hangul(GetStandard(deadline)), timer.Detail(GetAccurate(deadline))}
}
