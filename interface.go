package times

//Gadget is an interface for tools that uses time.
//Each Gadget has minimum unit as time.Unit, and It has times.TimeFormatter to represent current times as string.
//Examples of Gadget can be clock, timer, stopwatch.
type Gadget interface {
	Add(action func(current string))                 //Add adds function that would be called when duration of minimum unit has passed. Parameter current in function is string-conversion of current time that inner function can uses.
	AddAlarm(action func(current string), when Time) //AddAlarm adds function that would be called when Gadget reaches in selected time. Parameter current in function is string-conversion of current time that inner function can uses.
	Start()                                          //Start runs Gadget.
	Stop() string                                    //Stop calls Pause and Reset. It also returns string as current.
	Reset()                                          //Reset sets Gadget to the state when it firstly initialized.
	Pause()                                          //Pause stops ticker.
	Present() string                                 //Present returns current time.
}

//TimeFormatter is an interface that converts Time into string.
//TimeFormatter is used in times.Gadget for the use of showing times.
type TimeFormatter interface {
	Format(t Time) string //Format converts Time into string. This function can be used in structs that implement gadget.Gadget
}

//Time is interface that indicate specific time.
//Each Time should have time data and minimum Unit.
type Time interface {
	Tick()                   //Tick is called when time passed. This function add minimum Unit of this object.
	Rewind()                 //Rewind is called when time has been rewound. This function subtract minimum Unit of this object.
	MilliSecond() int        //MilliSecond returns millisecond of this object.
	Second() int             //Second returns second of this object.
	Minute() int             //Minute returns minute of this object.
	Hour() int               //Hour returns hour of this object.
	Day() int                //Day returns day of this object.
	Equal(to Time) bool      //Equal returns true when time of this object equals another one.
	SetMilliSecond(int) Time //SetMilliSecond sets millisecond of this object. This function also returns object itself.
	SetSecond(int) Time      //SetSecond sets millisecond of this object. This function also returns object itself.
	SetMinute(int) Time      //SetMinute sets millisecond of this object. This function also returns object itself.
	SetHour(int) Time        //SetHour sets millisecond of this object. This function also returns object itself.
	SetDay(int) Time         //SetDay sets millisecond of this object. This function also returns object itself.
	Serialize() string       //Serialize returns time data by string. This function is used for specifying time.
}
