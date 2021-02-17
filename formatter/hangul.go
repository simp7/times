package formatter

import (
	"fmt"
	"github.com/simp7/times"
	"strings"
)

type hangulFormatter struct {
	unit    []string
	isClock bool
}

//Hangul returns one of struct that implements times.TimeFormatter.
//Hangul shows time like 0분00초, And It can express time unit from seconds to day.
func Hangul() times.TimeFormatter {
	f := new(hangulFormatter)
	f.unit = []string{"일", "시간", "분", "초"}
	return f
}

func (f *hangulFormatter) Format(t times.Time) (result string) {

	addResult := func(data int, unit string) {
		if data != 0 {
			result += format(data, unit)
		}
	}

	data := []int{t.Day(), t.Hour(), t.Minute(), t.Second()}

	for i := range data {
		addResult(data[i], f.unit[i])
	}

	if result == "" {
		result = "0초" //zero-value
	}
	result = strings.TrimSpace(result)

	return

}

func format(data int, unit string) string {
	return fmt.Sprintf("%d%s ", data, unit)
}
