package action

import "github.com/simp7/times/model/tobject"

type Actions interface {
	Add(Action, tobject.Time)
	ActionsWhen(tobject.Time) Action
}

type actions struct {
	data           map[string]Action
	alwaysFunction Action
}

func NewActions() Actions {
	a := new(actions)
	a.data = make(map[string]Action)
	a.alwaysFunction = EmptyAction
	return a
}

func (a *actions) Add(action Action, time tobject.Time) {
	if time == nil {
		a.alwaysFunction = a.alwaysFunction.Add(action.Do)
	} else {
		a.set(time, action)
	}
}

func (a *actions) ActionsWhen(time tobject.Time) Action {
	defer a.delete(time)
	return a.alwaysFunction.Add(a.get(time).Do)
}

func (a *actions) get(time tobject.Time) Action {
	result := a.data[time.Serialize()]
	if result == nil {
		result = EmptyAction
	}
	return result
}

func (a *actions) set(time tobject.Time, action Action) {
	tmp := a.get(time)
	a.data[time.Serialize()] = tmp.Add(action.Do)
}

func (a *actions) delete(t tobject.Time) {
	delete(a.data, t.Serialize())
}
