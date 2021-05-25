package stopwatch

import (
	"github.com/simp7/times/formatter"
	"time"
)

//Standard returns gadget.Stopwatch that implements timeObject.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns gadget.Stopwatch that implements timeObject.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns gadget.Stopwatch that implements timeObject.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = New(time.Second, formatter.Standard())
	Hangul   = New(time.Second, formatter.Hangul())
	Detail   = New(time.Millisecond, formatter.Detail())
)
