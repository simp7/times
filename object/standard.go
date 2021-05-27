package object

import (
	"fmt"
	"github.com/simp7/times"
	"time"
)

type standard struct {
	second, minute, hour int8
	day                  int
}

//Standard is function that returns the struct that implements times.Object.
//Minimum unit of Standard is second.
func Standard(second, minute, hour, day int) times.Object {

	t := new(standard)

	t.SetSecond(second)
	t.SetMinute(minute)
	t.SetHour(hour)
	t.SetDay(day)

	return t

}

// StandardFor is function that gets built-in time.Time object and convert it to times.Object object.
// The other feature of StandardFor is same as Standard.
func StandardFor(t time.Time) times.Object {
	return Standard(t.Second(), t.Minute(), t.Hour(), t.Day())
}

//StandardZero is zero value of Object by using Standard.
func StandardZero() times.Object {
	return Standard(0, 0, 0, 0)
}

//Should be called after calculation
func (t *standard) trim() {

	t.trimIfAdded()
	t.trimIfSubtracted()

	if t.Day() < 0 {
		panic(ErrNegativeTime)
	}

}

func (t *standard) trimIfAdded() {
	if t.second >= 60 {
		t.minute += t.second / 60
		t.second %= 60
	}
	if t.minute >= 60 {
		t.hour += t.minute / 60
		t.minute %= 60
	}
	if t.day >= 24 {
		t.day += int(t.hour / 24)
		t.hour %= 60
	}
}

func (t *standard) trimIfSubtracted() {
	if t.second < 0 {
		t.minute--
		t.second += 60
	}
	if t.minute < 0 {
		t.hour--
		t.minute += 60
	}
	if t.hour < 0 {
		t.day--
		t.hour += 24
	}
}

func (t *standard) Tick() {
	t.second++
	t.trim()
}

func (t *standard) Rewind() {
	t.second--
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

func (t *standard) Day() int {
	return t.day
}

func (t *standard) Equal(another times.Object) bool {
	return t.Day() == another.Day() && t.Hour() == another.Hour() && t.Minute() == another.Minute() && t.Second() == another.Second() && t.MilliSecond() == another.MilliSecond()
}

func (t *standard) SetMilliSecond(int) times.Object {
	return t
}

func (t *standard) SetSecond(second int) times.Object {

	if second >= 60 || second < 0 {
		second = 0
	}

	t.second = int8(second)

	return t

}

func (t *standard) SetMinute(minute int) times.Object {

	if minute >= 60 || minute < 0 {
		minute = 0
	}

	t.minute = int8(minute)

	return t

}

func (t *standard) SetHour(hour int) times.Object {

	if hour >= 24 || hour < 0 {
		hour = 0
	}

	t.hour = int8(hour)

	return t

}

func (t *standard) SetDay(day int) times.Object {

	if day < 0 {
		day = 0
	}

	t.day = day

	return t

}

func (t *standard) Serialize() string {
	return fmt.Sprintf("%d/%d/%d/%d/%d", t.Day(), t.Hour(), t.Minute(), t.Second(), t.MilliSecond())
}
