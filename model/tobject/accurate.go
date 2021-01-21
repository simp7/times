package tobject

type accurate struct {
	Time
	ms int
}

func Accurate(millisecond int, second, minute, hour int8, day int) Time {

	a := new(accurate)
	a.Time= Standard(second, minute, hour, day)
	a.ms = millisecond
	a.trim()

	return a

}

func (t *accurate) Tick() {
	t.ms ++
	t.trim()
}

func (t *accurate) Rewind() {
	t.ms --
	t.trim()
}

func (t *accurate) trim() {

	if t.MilliSecond() >= 1000 {
		for i := 0; i < t.MilliSecond()/1000; i ++ {
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