package timer

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

var (
	Standard = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Standard(), t) }
	Hangul   = func(t tobject.Time) Timer { return New(tobject.Sec, formatter.Hangul(), t) }
	Detail   = func(t tobject.Time) Timer { return New(tobject.Ms, formatter.Detail(), t) }
)
