package action

import (
	"github.com/simp7/times"
)

type action struct {
	body func(string)
}

//NewAction capsulize function to action.
func NewAction(f func(string)) times.Action {
	var a = action{f}
	return a
}

func (a action) Add(action func(string)) times.Action {
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
