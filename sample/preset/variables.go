package preset

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/time"
)

//GetStandard returns time.Standard object that are set to seconds by parameter sec.
//Notice: GetStandard is just for sample. We recommend you to use time.Standard function as a substitute.
func GetStandard(sec int) times.Time {
	return time.StandardZero().SetSecond(sec)
}

//GetAccurate returns time.Accurate object that are set to seconds by parameter sec.
//Notice: GetAccurate is just for sample. We recommend you to use time.Accurate function as a substitute.
func GetAccurate(sec int) times.Time {
	return time.AccurateZero().SetSecond(sec)
}

//Stopwatches returns an array of gadget.Stopwatch.
//Notice: Stopwatches is just for sample.
func Stopwatches() []gadget.Stopwatch {
	return []gadget.Stopwatch{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail}
}

//Timers returns an array of gadget.Timer that seconds are set to value of parameter deadline.
//Notice: Timers is just for sample.
func Timers(deadline int) []gadget.Timer {
	return []gadget.Timer{timer.Standard(GetStandard(deadline)), timer.Hangul(GetStandard(deadline)), timer.Detail(GetAccurate(deadline))}
}
