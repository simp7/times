package times

//Formatter is an interface that converts Object into string.
//Formatter is used in times.Gadget for the use of showing times.
type Formatter interface {
	Format(t Object) string //Format converts Object into string. This function can be used in structs that implement gadget.Gadget
}
