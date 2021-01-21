package tobject

type Time interface {
	trim()
	Tick()
	Rewind()
	MilliSecond() int
	Second() int
	Minute() int
	Hour() int
	Day() int
	Equal(Time) bool
	SetMilliSecond(int) Time
	SetSecond(int) Time
	SetMinute(int) Time
	SetHour(int) Time
	SetDay(int) Time
}
