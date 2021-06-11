package stopwatch

import (
	"github.com/simp7/times/format"
	"time"
)

//Standard returns gadget.Stopwatch that implements object.Standard and format.Standard adopting minimum unit as second.
//Hangul returns gadget.Stopwatch that implements object.Standard and format.Hangul adopting minimum unit as second.
//Detail returns gadget.Stopwatch that implements object.Accurate and format.Detail adopting minimum unit as millisecond.
var (
	Standard = New(time.Second, format.Standard)
	Hangul   = New(time.Second, format.Hangul)
	Detail   = New(time.Millisecond, format.Detail)
)
