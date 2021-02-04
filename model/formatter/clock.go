package formatter

import (
	"fmt"
	"github.com/simp7/times/model/tobject"
)

type clockFormatter struct {
	is24 bool
}

func Clock(is24 bool) TimeFormatter {
	f := new(clockFormatter)
	f.is24 = is24
	return f
}

func (c *clockFormatter) Format(t tobject.Time) string {

	result := fmt.Sprintf("%s:%s", doubleDigitFormat(t.Minute()), doubleDigitFormat(t.Second()))

	if c.is24 {
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
