package action

//Action is kind of function that would be used in gadget.Gadget.
type Action interface {
	Add(func(string)) Action //Add returns action that has current function of this object and parameter function.
	Do(string)               //Do executes function that current object has.
}

type action struct {
	body func(string)
}

//NewAction capsulize function to action.
func NewAction(f func(string)) Action {
	var a = action{f}
	return a
}

func (a action) Add(action func(string)) Action {
	tmp := a.body
	a.body = func(s string) {
		tmp(s)
		action(s)
	}
	return a
}

func (a action) Do(s string) {
	a.body(s)
}
