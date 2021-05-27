package object

import (
	"fmt"
	"github.com/simp7/times"
	"time"
)

type accurate struct {
	times.Object
	ms int
}

//Accurate is function that returns the struct that implements times.Object.
//Minimum unit of Accurate is second.
//As Accurate returns the most-accurate times.Object object, It is encouraged to compare other object that implements times.Object with this.
func Accurate(millisecond, second, minute, hour, day int) times.Object {

	a := new(accurate)

	a.Object = Standard(second, minute, hour, day)
	a.SetMilliSecond(millisecond)

	return a

}

// AccurateFor is function that gets built-in time.Time object and convert it to the struct that implements times.Object. object.
// The other feature of AccurateFor is same as Accurate.
func AccurateFor(t time.Time) times.Object {
	return Accurate(t.Nanosecond()/1000000, t.Second(), t.Minute(), t.Hour(), t.Day())
}

//AccurateZero is zero value of Object by using Accurate.
func AccurateZero() times.Object {
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
			t.Object.Tick()
		}
		t.ms %= 1000
	}

	if t.MilliSecond() < 0 {
		t.Object.Rewind()
		t.ms += 1000
	}

}

func (t *accurate) MilliSecond() int {
	return t.ms
}

func (t *accurate) Equal(another times.Object) bool {
	return t.Day() == another.Day() && t.Hour() == another.Hour() && t.Minute() == another.Minute() && t.Second() == another.Second() && t.MilliSecond() == another.MilliSecond()
}

func (t *accurate) SetMilliSecond(ms int) times.Object {

	if t.ms >= 1000 || t.ms < 0 {
		t.ms = 0
	}

	t.ms = ms

	return t

}

func (t *accurate) SetSecond(second int) times.Object {
	t.Object.SetSecond(second)
	return t
}

func (t *accurate) SetMinute(minute int) times.Object {
	t.Object.SetMinute(minute)
	return t
}

func (t *accurate) SetHour(hour int) times.Object {
	t.Object.SetHour(hour)
	return t
}

func (t *accurate) SetDay(day int) times.Object {
	t.Object.SetDay(day)
	return t
}

func (t *accurate) Serialize() string {
	return fmt.Sprintf("%d/%d/%d/%d/%d", t.Day(), t.Hour(), t.Minute(), t.Second(), t.MilliSecond())
}
