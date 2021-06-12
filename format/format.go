package format

import (
	"fmt"
)

func doubleDigitFormat(a int) string {
	if a < 10 {
		return fmt.Sprintf("0%d", a)
	}
	return fmt.Sprintf("%d", a)
}

func tripleDigitFormat(a int) string {
	if a < 10 {
		return fmt.Sprintf("00%d", a)
	} else if a < 100 {
		return fmt.Sprintf("0%d", a)
	}
	return fmt.Sprintf("%d", a)
}
