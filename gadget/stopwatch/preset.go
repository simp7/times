package stopwatch

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

var (
	Standard = New(tobject.Sec, formatter.Standard())
	Hangul   = New(tobject.Sec, formatter.Hangul())
	Detail   = New(tobject.Ms, formatter.Detail())
)
