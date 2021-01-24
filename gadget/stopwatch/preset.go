package stopwatch

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

var (
	Standard = New(tobject.Sec, formatter.Standard()) //Standard returns Stopwatch that implements tobject.Standard and formatter.Standard adopting minimum unit as second.
	Hangul   = New(tobject.Sec, formatter.Hangul())   //Hangul returns Stopwatch that implements tobject.Standard and formatter.Hangul adopting minimum unit as second.
	Detail   = New(tobject.Ms, formatter.Detail())    //Detail returns Stopwatch that implements tobject.Accurate and formatter.Detail adopting minimum unit as millisecond.
)
