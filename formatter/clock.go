package formatter

import (
	"fmt"
	"github.com/simp7/times"
)

type clockFormatter struct {
	notation12 bool
}

//Clock returns one of struct that implements times.Formatter.
//Clock shows time like 11:23 PM or 23:23, And It is up to notation12 parameter.
func Clock(notation12 bool) times.Formatter {
	f := new(clockFormatter)
	f.notation12 = notation12
	return f
}

func (c *clockFormatter) Format(t times.Object) string {

	result := fmt.Sprintf("%s:%s", doubleDigitFormat(t.Minute()), doubleDigitFormat(t.Second()))

	if c.notation12 {
		hour, suffix := decomposeHour(t.Hour())
		return fmt.Sprintf("%d:%s %s", hour, result, suffix)
	}
	return fmt.Sprintf("%d:%s", t.Hour(), result)

}

func decomposeHour(hour int) (result int, suffix string) {
	result = hour
	suffix = "AM"
	if hour > 12 {
		result -= 12
		suffix = "PM"
	}
	return
}
