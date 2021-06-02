package action

import (
	"github.com/simp7/times"
)

type action func(string)

//NewAction capsulize function to action.
func NewAction(f func(string)) times.Action {
	return action(f)
}

func (a action) Add(action times.Action) times.Action {
	tmp := a
	return NewAction(func(s string) {
		tmp(s)
		action.Do(s)
	})
}

func (a action) Do(s string) {
	a(s)
}
