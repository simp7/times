package timer

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/object"
	"sync"
	"time"
)

type timer struct {
	ticker    *gadget.Ticker
	present   times.Object
	deadline  times.Object
	format    times.Format
	unit      time.Duration
	once      sync.Once
	isRunning bool
	actions   times.Actions
}

//New returns struct that implements times.Gadget.
//parameter unit is for ticking rate, format for formatting time to string, and deadline for deadline of timer.
func New(unit time.Duration, format times.Format, deadline times.Object) times.Gadget {

	t := new(timer)

	t.unit = unit
	t.format = format
	t.deadline = deadline
	t.isRunning = false

	t.ticker = gadget.NewTicker(unit)

	t.Reset()

	return t

}

func (t *timer) Start() {
	t.isRunning = true
	t.work()
}

func (t *timer) getAction() times.Action {
	return t.actions.ActionsWhen(t.present)
}

func (t *timer) do() {
	t.getAction().Do(t.present)
}

func (t *timer) work() {
	t.ticker.Start(func() {
		t.present.Rewind()
		t.once.Do(t.present.Tick)
		t.do()
	})
}

func (t *timer) Stop() string {

	result := t.Present()

	t.Pause()
	t.Reset()

	return result

}

func (t *timer) Add(action times.Action) {
	t.actions.Add(action, nil)
}

func (t *timer) AddAlarm(action times.Action, when times.Object) {
	t.actions.Add(action, when)
}

func (t *timer) Reset() {

	preset := t.deadline
	if t.unit == time.Millisecond {
		t.present = object.AccurateZero()
	} else {
		t.present = object.StandardZero()
	}

	t.present.SetMilliSecond(preset.MilliSecond()).
		SetSecond(preset.Second()).
		SetMinute(preset.Minute()).
		SetHour(preset.Hour()).
		SetDay(preset.Day())

	t.actions = action.NewActions()
	t.AddAlarm(action.NewAction(func(times.Object) { t.Stop() }), object.StandardZero())

}

func (t *timer) resetPresent() {

	var result times.Object
	preset := t.deadline

	if t.unit == time.Millisecond {
		result = object.AccurateZero()
	} else {
		result = object.StandardZero()
	}

	result.SetMilliSecond(preset.MilliSecond())
	result.SetSecond(preset.Second())
	result.SetMinute(preset.Minute())
	result.SetHour(preset.Hour())
	result.SetDay(preset.Day())

	t.present = result

}

func (t *timer) Pause() {
	if t.isRunning {
		t.ticker.Stop()
		t.isRunning = false
		t.once = sync.Once{}
	}
}

func (t *timer) Present() string {
	return t.format(t.present)
}

func (t *timer) GetFormat() times.Format {
	return t.format
}
