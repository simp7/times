package preset

import (
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/model/tobject"
)

func GetStandard(sec int) tobject.Time {
	return tobject.StandardZero().SetSecond(sec)
}

func GetAccurate(sec int) tobject.Time {
	return tobject.AccurateZero().SetSecond(sec)
}

func Stopwatches() []stopwatch.Stopwatch {
	return []stopwatch.Stopwatch{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail}
}

func Timers(deadline int) []timer.Timer {
	return []timer.Timer{timer.Standard(GetStandard(deadline)), timer.Hangul(GetStandard(deadline)), timer.Detail(GetAccurate(deadline))}
}
