package formatter

import (
	"github.com/simp7/times"
)

//Detail is a function that implements times.Format.
//Detail shows time like 0:00:000, And It can express time unit from millisecond to day.
func Detail(t times.Object) string {
	return Standard(t) + ":" + tripleDigitFormat(t.MilliSecond())
}
