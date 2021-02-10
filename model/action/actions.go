package action

import "github.com/simp7/times/model/tobject"

//Actions is an interface that include a set of Action.
//Actions would be used in gadget.Gadget.
type Actions interface {
	Add(Action, tobject.Time)        //Add adds Action to Actions in designated time.
	ActionsWhen(tobject.Time) Action //ActionsWhen returns Action in designated time.
}

type actions struct {
	data           map[string]Action
	alwaysFunction Action
}

//NewActions Returns Empty Actions.
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
