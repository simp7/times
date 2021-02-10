package preset

import (
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/model/tobject"
)

//Notice: GetStandard is just for sample. We recommend you to use tobject.Standard function.
//GetStandard returns tobject.Standard object that are set to seconds by parameter sec.
func GetStandard(sec int) tobject.Time {
	return tobject.StandardZero().SetSecond(sec)
}

//Notice: GetAccurate is just for sample. We recommend you to use tobject.Accurate function.
//GetAccurate returns tobject.Accuratre object that are set to seconds by parameter sec.
func GetAccurate(sec int) tobject.Time {
	return tobject.AccurateZero().SetSecond(sec)
}

//Notice: Stopwatches is just for sample.
//Stopwatches returns stopwatches.
func Stopwatches() []stopwatch.Stopwatch {
	return []stopwatch.Stopwatch{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail}
}

//Notice: Timers is just for sample.
//Timers returns Timers that seconds are set to value of parameter deadline.
func Timers(deadline int) []timer.Timer {
	return []timer.Timer{timer.Standard(GetStandard(deadline)), timer.Hangul(GetStandard(deadline)), timer.Detail(GetAccurate(deadline))}
}
