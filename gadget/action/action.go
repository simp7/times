package action

import (
	"github.com/simp7/times"
)

type action func(object times.Object)

//NewAction capsulize function to action.
func NewAction(f func(object times.Object)) times.Action {
	return action(f)
}

func (a action) Add(action times.Action) times.Action {
	tmp := a
	return NewAction(func(o times.Object) {
		tmp(o)
		action.Do(o)
	})
}

func (a action) Do(object times.Object) {
	a(object)
}
