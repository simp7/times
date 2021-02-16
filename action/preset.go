package action

//EmptyAction is zero-value of action.
var (
	EmptyAction = NewAction(func(string) {})
)
