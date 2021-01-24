package formatter

import (
	"fmt"
	"github.com/simp7/times/model/tobject"
)

type standardFormatter struct {
}

//Standard returns one of struct that implements TimeFormatter.
//Standard shows time like 0:00, And It can express time unit from second to day.
func Standard() TimeFormatter {
	f := new(standardFormatter)
	return f
}

func (f *standardFormatter) Format(t tobject.Time) string {
	if t.Second() < 10 {
		return fmt.Sprintf("%d:0%d", t.Minute(), t.Second())
	} else if t.Hour() == 0 {
		return fmt.Sprintf("%d:%d", t.Minute(), t.Second())
	} else if t.Day() == 0 {
		return fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
	} else {
		return fmt.Sprintf("%d:%d:%d:%d", t.Day(), t.Hour(), t.Minute(), t.Second())
	}
}
