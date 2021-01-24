package formatter

import (
	"fmt"
	"github.com/simp7/times/model/tobject"
)

type detailFormatter struct {
	standardFormatter
}

//Detail returns one of struct that implements TimeFormatter.
//Detail shows time like 0:00:000, And It can express time unit from millisecond to day.
func Detail() TimeFormatter {
	t := new(detailFormatter)
	return t
}

func (f *detailFormatter) Format(t tobject.Time) string {
	result := f.standardFormatter.Format(t)
	if t.MilliSecond() < 10 {
		result += fmt.Sprintf(":00%d", t.MilliSecond())
	} else if t.MilliSecond() < 100 {
		result += fmt.Sprintf(":0%d", t.MilliSecond())
	} else {
		result += fmt.Sprintf(":%d", t.MilliSecond())
	}
	return result
}
