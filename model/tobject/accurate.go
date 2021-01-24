package tobject

type accurate struct {
	Time
	ms int
}

//Accurate is function that returns Time object.
//Minimum unit of Accurate is second.
//As this function returns the most-accurate Time object, It is encouraged to compare Time object with this.
func Accurate(millisecond int, second, minute, hour int8, day int) Time {

	a := new(accurate)
	a.Time = Standard(second, minute, hour, day)
	a.ms = millisecond
	a.trim()

	return a

}

//AccurateZero is zero value of Time by using Accurate.
func AccurateZero() Time {
	return Accurate(0, 0, 0, 0, 0)
}

func (t *accurate) Tick() {
	t.ms++
	t.trim()
}

func (t *accurate) Rewind() {
	t.ms--
	t.trim()
}

//Should be called after calculation
func (t *accurate) trim() {

	if t.MilliSecond() >= 1000 {
		for i := 0; i < t.MilliSecond()/1000; i++ {
			t.Time.Tick()
		}
		t.ms %= 1000
	}

	if t.MilliSecond() < 0 {
		t.Time.Rewind()
		t.ms += 1000
	}

}

func (t *accurate) MilliSecond() int {
	return t.ms
}

func (t *accurate) Equal(another Time) bool {
	return t.Day() == another.Day() && t.Hour() == another.Hour() && t.Minute() == another.Minute() && t.Second() == another.Second() && t.MilliSecond() == another.MilliSecond()
}

func (t *accurate) SetMilliSecond(ms int) Time {
	t.ms = ms
	return t
}

func (t *accurate) SetSecond(second int) Time {
	t.Time.SetSecond(second)
	return t
}

func (t *accurate) SetMinute(minute int) Time {
	t.Time.SetMinute(minute)
	return t
}

func (t *accurate) SetHour(hour int) Time {
	t.Time.SetHour(hour)
	return t
}

func (t *accurate) SetDay(day int) Time {
	t.Time.SetDay(day)
	return t
}
