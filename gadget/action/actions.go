package action

import (
	"github.com/simp7/times"
)

type actions struct {
	data           map[string]times.Action
	alwaysFunction times.Action
}

//NewActions Returns Empty Actions.
func NewActions() times.Actions {
	a := new(actions)
	a.data = make(map[string]times.Action)
	a.alwaysFunction = EmptyAction
	return a
}

func (a *actions) Add(action times.Action, time times.Object) {
	if time == nil {
		a.alwaysFunction = a.alwaysFunction.Add(action)
	} else {
		a.set(time, action)
	}
}

func (a *actions) ActionsWhen(time times.Object) times.Action {
	defer a.delete(time)
	return a.alwaysFunction.Add(a.get(time))
}

func (a *actions) get(time times.Object) times.Action {
	result := a.data[time.Serialize()]
	if result == nil {
		result = EmptyAction
	}
	return result
}

func (a *actions) set(time times.Object, action times.Action) {
	tmp := a.get(time)
	a.data[time.Serialize()] = tmp.Add(action)
}

func (a *actions) delete(t times.Object) {
	delete(a.data, t.Serialize())
}
