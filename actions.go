package times

//Actions is an interface that includes a set of Action.
//Actions would be used in times.Gadget.
type Actions interface {
	Add(Action, Object)        //Add adds Action to Actions in designated time.
	ActionsWhen(Object) Action //ActionsWhen returns Action in designated time.
}
