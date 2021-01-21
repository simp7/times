package tobject

type standard struct {
	second, minute, hour int8
	day int
}

func Standard(second, minute, hour int8, day int) Time {

	t := new(standard)
	t.second, t.minute, t.hour, t.day = second, minute, hour, day
	t.trim()
	return t

}


//Should be called after calculation
func (t *standard) trim() {
	t.trimIfAdded()
	t.trimIfSubtracted()
}

func (t *standard) trimIfAdded() {
	if t.second >= 60 {
		t.minute += t.second/60
		t.second %= 60
	}
	if t.minute >= 60 {
		t.hour += t.minute/60
		t.minute %= 60
	}
	if t.day >= 24 {
		t.day += int(t.hour/24)
		t.hour %= 60
	}
}

func (t *standard) trimIfSubtracted() {
	if t.second < 0 {
		t.minute --
		t.second += 60
	}
	if t.minute < 0 {
		t.hour --
		t.minute += 60
	}
	if t.hour < 0 {
		t.day --
		t.hour += 24
	}
}

func (t *standard) Tick() {
	t.second ++
	t.trim()
}

func (t *standard) Rewind() {
	t.second --
	t.trim()
}

func (t *standard) MilliSecond() int {
	return 0
}

func (t *standard) Second() int {
	return int(t.second)
}

func (t *standard) Minute() int {
	return int(t.minute)
}

func (t *standard) Hour() int {
	return int(t.hour)
}

func (t *standard) Day() int  {
	return t.day
}

func (t *standard) Equal(another Time) bool {
	return t.Day() == another.Day() && t.Hour() == another.Hour() && t.Minute() == another.Minute() && t.Second() == another.Second() && t.MilliSecond() == another.MilliSecond()
}