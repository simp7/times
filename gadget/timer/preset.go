package timer

import (
	"github.com/simp7/times"
	"github.com/simp7/times/formatter"
	"time"
)

//Standard returns gadget.Timer that implements timeobject.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns gadget.Timer that implements timeobject.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns gadget.Timer that implements timeobject.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = func(t times.Time) times.Gadget { return New(time.Second, formatter.Standard(), t) }
	Hangul   = func(t times.Time) times.Gadget { return New(time.Second, formatter.Hangul(), t) }
	Detail   = func(t times.Time) times.Gadget { return New(time.Millisecond, formatter.Detail(), t) }
)
