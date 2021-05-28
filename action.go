package times

//Action is an interface that includes function.
//Action would be used in Actions and times.Gadget.
type Action interface {
	Add(func(string)) Action //Add returns action that has current function of this object and parameter function.
	Do(string)               //Do executes function that current object has.
}
