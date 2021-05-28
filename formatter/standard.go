package formatter

import (
	"fmt"
	"github.com/simp7/times"
)

type standardFormatter struct {
}

//Standard returns one of struct that implements times.Formatter.
//Standard shows time like 0:00, And It can express time unit from second to day.
func Standard() times.Formatter {
	return new(standardFormatter)
}

func (f *standardFormatter) Format(t times.Object) string {

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
