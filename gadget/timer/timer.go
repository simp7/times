package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
)

type timer struct {
	ticker    gadget.Ticker
	present   tobject.Time
	deadline  tobject.Time
	formatter formatter.TimeFormatter
	unit      tobject.Unit
	once      sync.Once
	isRunning bool
	actions   action.Actions
}

//New returns struct that implements Timer.
//parameter unit is for ticking rate, formatter for formatting time to string, and deadline for deadline of timer.
func New(unit tobject.Unit, formatter formatter.TimeFormatter, deadline tobject.Time) gadget.Timer {

	t := new(timer)

	t.unit = unit
	t.formatter = formatter
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

func (t *timer) getAction() action.Action {
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

func (t *timer) AddAlarm(f func(string), when tobject.Time) {
	t.actions.Add(action.NewAction(f), when)
}

func (t *timer) Reset() {

	preset := t.deadline
	if t.unit == tobject.Ms {
		t.present = tobject.AccurateZero()
	} else {
		t.present = tobject.StandardZero()
	}

	t.present.SetMilliSecond(preset.MilliSecond()).
		SetSecond(preset.Second()).
		SetMinute(preset.Minute()).
		SetHour(preset.Hour()).
		SetDay(preset.Day())

	t.actions = action.NewActions()
	t.AddAlarm(func(string) { t.Stop() }, tobject.StandardZero())

}

func (t *timer) resetPresent() {

	var result tobject.Time
	preset := t.deadline

	if t.unit == tobject.Ms {
		result = tobject.AccurateZero()
	} else {
		result = tobject.StandardZero()
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
