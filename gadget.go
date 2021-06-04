package times

//Gadget is an interface for tools that uses time.
//Each Gadget has minimum unit as time.Unit, and It has times.Format to represent current times as string.
//Examples of Gadget can be clock, timer, stopwatch.
type Gadget interface {
	Add(action Action)                   //Add adds function that would be called when duration of minimum unit has passed. Parameter current in function is string-conversion of current time that inner function can uses.
	AddAlarm(action Action, when Object) //AddAlarm adds function that would be called when Gadget reaches in selected time. Parameter current in function is string-conversion of current time that inner function can uses.
	Start()                              //Start runs Gadget.
	Stop() string                        //Stop calls Pause and Reset. It also returns string as current.
	Reset()                              //Reset sets Gadget to the state when it firstly initialized.
	Pause()                              //Pause stops ticker.
	Present() string                     //Present returns current time.
	GetFormat() Format                   //GetFormat returns inner format of the Gadget.
}
