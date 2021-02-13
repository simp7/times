package preset

import (
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/model/tobject"
)

//GetStandard returns tobject.Standard object that are set to seconds by parameter sec.
//Notice: GetStandard is just for sample. We recommend you to use tobject.Standard function as a substitute.
func GetStandard(sec int) tobject.Time {
	return tobject.StandardZero().SetSecond(sec)
}

//GetAccurate returns tobject.Accuratre object that are set to seconds by parameter sec.
//Notice: GetAccurate is just for sample. We recommend you to use tobject.Accurate function as a substitute.
func GetAccurate(sec int) tobject.Time {
	return tobject.AccurateZero().SetSecond(sec)
}

//Stopwatches returns an array of stopwatch.Stopwatch.
//Notice: Stopwatches is just for sample.
func Stopwatches() []stopwatch.Stopwatch {
	return []stopwatch.Stopwatch{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail}
}

//Timers returns an array of timer.Timer that seconds are set to value of parameter deadline.
//Notice: Timers is just for sample.
func Timers(deadline int) []timer.Timer {
	return []timer.Timer{timer.Standard(GetStandard(deadline)), timer.Hangul(GetStandard(deadline)), timer.Detail(GetAccurate(deadline))}
}
