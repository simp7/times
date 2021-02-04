package formatter

import (
	"github.com/simp7/times/model/tobject"
)

type detailFormatter struct {
	standardFormatter
}

//Detail returns one of struct that implements TimeFormatter.
//Detail shows time like 0:00:000, And It can express time unit from millisecond to day.
func Detail() TimeFormatter {
	return new(detailFormatter)
}

func (f *detailFormatter) Format(t tobject.Time) string {
	result := f.standardFormatter.Format(t)
	result = result + ":" + tripleDigitFormat(t.MilliSecond())
	return result
}
