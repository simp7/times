package clock

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

func Standard() Clock {
	return New(tobject.Sec, formatter.Clock(true))
}
