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

func (a *actions) Add(action times.Action, time times.Time) {
	if time == nil {
		a.alwaysFunction = a.alwaysFunction.Add(action.Do)
	} else {
		a.set(time, action)
	}
}

func (a *actions) ActionsWhen(time times.Time) times.Action {
	defer a.delete(time)
	return a.alwaysFunction.Add(a.get(time).Do)
}

func (a *actions) get(time times.Time) times.Action {
	result := a.data[time.Serialize()]
	if result == nil {
		result = EmptyAction
	}
	return result
}

func (a *actions) set(time times.Time, action times.Action) {
	tmp := a.get(time)
	a.data[time.Serialize()] = tmp.Add(action.Do)
}

func (a *actions) delete(t times.Time) {
	delete(a.data, t.Serialize())
}