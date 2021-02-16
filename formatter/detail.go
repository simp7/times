package formatter

import (
	"github.com/simp7/times"
)

type detailFormatter struct {
	standardFormatter
}

//Detail returns one of struct that implements times.TimeFormatter.
//Detail shows time like 0:00:000, And It can express time unit from millisecond to day.
func Detail() times.TimeFormatter {
	return new(detailFormatter)
}

func (f *detailFormatter) Format(t times.Time) string {
	result := f.standardFormatter.Format(t)
	result = result + ":" + tripleDigitFormat(t.MilliSecond())
	return result
}
