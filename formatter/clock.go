package formatter

import (
	"fmt"
	"github.com/simp7/times"
)

type clockFormatter struct {
}

//Clock returns one of struct that implements times.Format.
//Clock shows time like 11:23 PM or 23:23, And It is up to notation12 parameter.
var Clock = &clockFormatter{}

//Notation12 is a function that returns in form of format with notation of AM/PM of clock.
func (c *clockFormatter) Notation12(t times.Object) string {

	result := fmt.Sprintf("%s:%s", doubleDigitFormat(t.Minute()), doubleDigitFormat(t.Second()))
	hour, suffix := decomposeHour(t.Hour())

	return fmt.Sprintf("%d:%s %s", hour, result, suffix)

}

//Notation24 is a function that returns in form of format with notation of clock.
func (c *clockFormatter) Notation24(t times.Object) string {
	return fmt.Sprintf("%d:%s:%s", t.Hour(), doubleDigitFormat(t.Minute()), doubleDigitFormat(t.Second()))
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
