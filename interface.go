package times

import "github.com/simp7/times/model/tobject"

//Gadget is an interface for tools that uses time.
//Each Gadget has minimum unit as tobject.Unit, and It has formatter.TimeFormatter to represent current times as string.
//Examples of Gadget can be clock, timer, stopwatch.
type Gadget interface {
	Add(action func(current string))                         //Add adds function that would be called when duration of minimum unit has passed. Parameter @current in function is string-conversion of current time that inner function can uses.
	AddAlarm(action func(current string), when tobject.Time) //AddAlarm adds function that would be called when Gadget reaches in selected time. Parameter @current in function is string-conversion of current time that inner function can uses.
	Start()                                                  //Start runs Gadget.
	Stop() string                                            //Stop calls Pause and Reset. It also returns string as current.
	Reset()                                                  //Reset sets Gadget to the state when it firstly initialized.
	Pause()                                                  //Pause stops ticker.
	Present() string                                         //Present returns current time.
}
