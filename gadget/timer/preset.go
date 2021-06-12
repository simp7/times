package timer

import (
	"github.com/simp7/times"
	"github.com/simp7/times/format"
	"time"
)

//Standard returns gadget.Timer that implements object.Standard and format.Standard adopting minimum unit as second.
//Hangul returns gadget.Timer that implements object.Standard and format.Hangul adopting minimum unit as second.
//Detail returns gadget.Timer that implements object.Accurate and format.Detail adopting minimum unit as millisecond.
var (
	Standard = func(t times.Object) times.Gadget { return New(time.Second, format.Standard, t) }
	Hangul   = func(t times.Object) times.Gadget { return New(time.Second, format.Hangul, t) }
	Detail   = func(t times.Object) times.Gadget { return New(time.Millisecond, format.Detail, t) }
)
