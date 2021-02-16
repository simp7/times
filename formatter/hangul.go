package formatter

import (
	"fmt"
	"github.com/simp7/times"
)

type hangulFormatter struct {
}

//Hangul returns one of struct that implements times.TimeFormatter.
//Hangul shows time like 0분00초, And It can express time unit from seconds to day.
func Hangul() times.TimeFormatter {
	return new(hangulFormatter)
}

func (f *hangulFormatter) Format(t times.Time) (result string) {

	switch {
	case t.Day() != 0:
		result += fmt.Sprintf("%d일 ", t.Day())
		fallthrough
	case t.Hour() != 0:
		result += fmt.Sprintf("%d시 ", t.Hour())
		fallthrough
	case t.Minute() != 0:
		result += fmt.Sprintf("%d분 ", t.Minute())
		fallthrough
	default:
		result += fmt.Sprintf("%d초", t.Second())
	}

	return

}
