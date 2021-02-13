package stopwatch

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

//Standard returns Stopwatch that implements tobject.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns Stopwatch that implements tobject.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns Stopwatch that implements tobject.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = New(tobject.Sec, formatter.Standard())
	Hangul   = New(tobject.Sec, formatter.Hangul())
	Detail   = New(tobject.Ms, formatter.Detail())
)
