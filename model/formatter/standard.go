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
	return new(standardFormatter)
}

func (f *standardFormatter) Format(t tobject.Time) string {

	sec := doubleDigitFormat(t.Second())
	min := doubleDigitFormat(t.Minute())
	hour := doubleDigitFormat(t.Hour())

	if t.Hour() == 0 {
		return fmt.Sprintf("%s:%s", min, sec)
	} else if t.Day() == 0 {
		return fmt.Sprintf("%s:%s:%s", hour, min, sec)
	} else {
		return fmt.Sprintf("%d:%s:%s:%s", t.Day(), hour, min, sec)
	}

}
