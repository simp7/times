package timer

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

//Standard returns Timer that implements tobject.Standard and formatter.Standard adopting minimum unit as second.
//Hangul returns Timer that implements tobject.Standard and formatter.Hangul adopting minimum unit as second.
//Detail returns Timer that implements tobject.Accurate and formatter.Detail adopting minimum unit as millisecond.
var (
	Standard = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Standard(), t) }
	Hangul   = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Hangul(), t) }
	Detail   = func(t tobject.Time) Timer { return New(tobject.Ms, formatter.Detail(), t) }
)
