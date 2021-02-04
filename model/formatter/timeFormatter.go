package formatter

import (
	"fmt"
	"github.com/simp7/times/model/tobject"
)

//TimeFormatter is an interface that converts tobject.Time into string.
//TimeFormatter is used in gadgets for the use of showing times.
type TimeFormatter interface {
	Format(t tobject.Time) string //Format converts tobject.Time into string. This function can be used in structs that implement gadget.Gadget
}

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
