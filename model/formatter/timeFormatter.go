package formatter

import (
	"github.com/simp7/times/model/tobject"
)

//TimeFormatter is an interface that converts tobject.Time into string.
//TimeFormatter is used in gadgets for the use of showing times.
type TimeFormatter interface {
	Format(t tobject.Time) string //Format converts tobject.Time into string. This function can be used in structs that implement gadget.Gadget
}
