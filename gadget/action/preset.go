package action

import "github.com/simp7/times"

//EmptyAction is zero-value of action.
var (
	EmptyAction = NewAction(func(object times.Object) {})
)
