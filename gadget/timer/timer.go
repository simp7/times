package timer

import (
	"github.com/simp7/times"
	"github.com/simp7/times/action"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/ticker"
	"github.com/simp7/times/ttime"
	"sync"
)

type timer struct {
	ticker    gadget.Ticker
	present   times.Time
	deadline  times.Time
	formatter times.TimeFormatter
	unit      times.Unit
	once      sync.Once
	isRunning bool
	actions   times.Actions
}

//New returns struct that implements gadget.Timer.
//parameter unit is for ticking rate, formatter for formatting time to string, and deadline for deadline of timer.
func New(unit times.Unit, formatter times.TimeFormatter, deadline times.Time) gadget.Timer {

	t := new(timer)

	t.unit = unit
	t.formatter = formatter
	t.deadline = deadline
	t.isRunning = false

	t.ticker = ticker.NewTicker(unit)

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
	current := t.Present()
	t.getAction().Do(current)
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

func (t *timer) Add(f func(string)) {
	t.actions.Add(action.NewAction(f), nil)
}

func (t *timer) AddAlarm(f func(string), when times.Time) {
	t.actions.Add(action.NewAction(f), when)
}

func (t *timer) Reset() {

	preset := t.deadline
	if t.unit == times.Ms {
		t.present = ttime.AccurateZero()
	} else {
		t.present = ttime.StandardZero()
	}

	t.present.SetMilliSecond(preset.MilliSecond()).
		SetSecond(preset.Second()).
		SetMinute(preset.Minute()).
		SetHour(preset.Hour()).
		SetDay(preset.Day())

	t.actions = action.NewActions()
	t.AddAlarm(func(string) { t.Stop() }, ttime.StandardZero())

}

func (t *timer) resetPresent() {

	var result times.Time
	preset := t.deadline

	if t.unit == times.Ms {
		result = ttime.AccurateZero()
	} else {
		result = ttime.StandardZero()
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
	return t.formatter.Format(t.present)
}
