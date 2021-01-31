package action

type Action interface {
	Add(func(string)) Action
	Do(string)
}

type action struct {
	body func(string)
}

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
