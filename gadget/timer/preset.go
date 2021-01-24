package timer

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

var (
	Standard = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Standard(), t) } //Standard returns Timer that implements tobject.Standard and formatter.Standard adopting minimum unit as second.
	Hangul   = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Hangul(), t) }   //Hangul returns Timer that implements tobject.Standard and formatter.Hangul adopting minimum unit as second.
	Detail   = func(t tobject.Time) Timer { return New(tobject.Ms, formatter.Detail(), t) }    //Detail returns Timer that implements tobject.Accurate and formatter.Detail adopting minimum unit as millisecond.
)
