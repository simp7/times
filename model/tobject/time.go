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
}