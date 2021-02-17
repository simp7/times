package stopwatch

import (
	"github.com/simp7/times"
	"github.com/simp7/times/formatter"
)

//Standard returns gadget.Stopwatch that implements ttime.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns gadget.Stopwatch that implements ttime.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns gadget.Stopwatch that implements ttime.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = New(times.Sec, formatter.Standard())
	Hangul   = New(times.Sec, formatter.Hangul())
	Detail   = New(times.Ms, formatter.Detail())
)
