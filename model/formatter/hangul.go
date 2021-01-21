package formatter

import (
	"fmt"
	"github.com/simp7/times/model/tobject"
)

type hangulFormatter struct {
}

func Hangul() TimeFormatter {
	f := new(hangulFormatter)
	return f
}

func (f *hangulFormatter) Format(t tobject.Time) (result string) {

	switch {
	case t.Day() != 0:
		result += fmt.Sprintf("%d일", t.Day())
		fallthrough
	case t.Hour() != 0:
		result += fmt.Sprintf("%d시", t.Hour())
		fallthrough
	case t.Minute() != 0:
		result += fmt.Sprintf("%d분", t.Minute())
		fallthrough
	default:
		result += fmt.Sprintf("%d초", t.Second())
	}
	return
}
