package stopwatch

import (
	"times/model/formatter"
	"times/model/tobject"
)

var (
	Standard = New(tobject.Sec, formatter.Standard())
	Hangul = New(tobject.Sec, formatter.Hangul())
	Detail = New(tobject.Ms, formatter.Detail())
)