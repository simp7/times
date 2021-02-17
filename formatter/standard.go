package formatter

import (
	"fmt"
	"github.com/simp7/times"
	"strconv"
)

type standardFormatter struct {
}

//Standard returns one of struct that implements times.TimeFormatter.
//Standard shows time like 0:00, And It can express time unit from second to day.
func Standard() times.TimeFormatter {
	return new(standardFormatter)
}

func (f *standardFormatter) Format(t times.Time) (result string) {

	addResult := func(data string) {
		result = fmt.Sprintf("%s:", data) + result
	}

	sec := doubleDigitFormat(t.Second())
	min := doubleDigitFormat(t.Minute())
	hour := doubleDigitFormat(t.Hour())
	day := strconv.Itoa(t.Day())

	result = fmt.Sprintf("%s:%s", min, sec)

	if t.Hour() == 0 && t.Day() == 0 {
		return
	}
	addResult(hour)
	if t.Day() != 0 {
		addResult(day)
	}

	return

}
