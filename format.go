package times

//Format is an interface that converts Object into string.
//Format is used in times.Gadget for the use of showing times.
type Format func(t Object) string
