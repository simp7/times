package object

import "github.com/simp7/times"

//Formatter is an interface that converts Object into string.
//Formatter is used in times.Gadget for the use of showing times.
type Formatter interface {
	Format(object times.Object) string //Format converts Object into string. This function can be used in structs that implement gadget.Gadget
}
