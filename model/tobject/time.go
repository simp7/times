package tobject

//Time is interface that indicate specific time.
//Each Time should have time data and minimum Unit.
type Time interface {
	trim()                   //trim should be called when inner time data has changed. This function would fix object to fit structure.
	Tick()                   //Tick is called when time passed. This function add minimum Unit of this object.
	Rewind()                 //Rewind is called when time has been rewound. This function subtract minimum Unit of this object.
	MilliSecond() int        //MilliSecond returns millisecond of this object.
	Second() int             // Second returns second of this object.
	Minute() int             // Minute returns minute of this object.
	Hour() int               //Hour returns hour of this object.
	Day() int                //Day returns day of this object.
	Equal(to Time) bool      //Equal returns true when time of this object equals another one.
	SetMilliSecond(int) Time //SetMilliSecond sets millisecond of this object. This function also returns object itself.
	SetSecond(int) Time      //SetSecond sets millisecond of this object. This function also returns object itself.
	SetMinute(int) Time      //SetMinute sets millisecond of this object. This function also returns object itself.
	SetHour(int) Time        //SetHour sets millisecond of this object. This function also returns object itself.
	SetDay(int) Time         //SetDay sets millisecond of this object. This function also returns object itself.
}
