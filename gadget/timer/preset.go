package timer

import (
	"github.com/simp7/times"
	"github.com/simp7/times/formatter"
	"github.com/simp7/times/gadget"
)

//Standard returns gadget.Timer that implements ttime.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns gadget.Timer that implements ttime.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns gadget.Timer that implements ttime.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = func(t times.Time) gadget.Timer { return New(times.Sec, formatter.Standard(), t) }
	Hangul   = func(t times.Time) gadget.Timer { return New(times.Sec, formatter.Hangul(), t) }
	Detail   = func(t times.Time) gadget.Timer { return New(times.Ms, formatter.Detail(), t) }
)
