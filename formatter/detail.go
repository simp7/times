package formatter

import (
	"github.com/simp7/times"
)

type detailFormatter struct {
	standardFormatter
}

//Detail returns one of struct that implements times.Formatter.
//Detail shows time like 0:00:000, And It can express time unit from millisecond to day.
func Detail() times.Formatter {
	return new(detailFormatter)
}

func (f *detailFormatter) Format(t times.Object) string {
	result := f.standardFormatter.Format(t)
	result = result + ":" + tripleDigitFormat(t.MilliSecond())
	return result
}
