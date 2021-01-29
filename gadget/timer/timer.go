package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

//Timer is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Timer interface {
	gadget.Gadget
	DoWhenFinished(func()) //DoWhenFinished is called when time of timer becomes zero.
}

type timer struct {
	ticker      gadget.Ticker
	present     tobject.Time
	deadline    tobject.Time
	formatter   formatter.TimeFormatter
	unit        tobject.Unit
	actions     []func(string)
	finalAction func()
}

func New(u tobject.Unit, f formatter.TimeFormatter, deadline tobject.Time) Timer {

	t := new(timer)

	t.unit = u
	t.formatter = f
	t.deadline = deadline

	t.ticker = gadget.NewTicker(u)

	t.Reset()

	return t

}

func (t *timer) Start() {
	t.work()
}

func (t *timer) do() {
	current := t.formatter.Format(t.present)
	for _, action := range t.actions {
		action(current)
	}
}

func (t *timer) work() {

	t.ticker.Start(func() {

		t.do()

		if t.present.Equal(tobject.AccurateZero()) {
			t.Stop()
			return
		}

		t.present.Rewind()

	})

}

func (t *timer) Stop() string {

	result := t.formatter.Format(t.present)

	t.finalAction()

	t.Pause()
	t.Reset()

	return result

}

func (t *timer) Add(action func(string)) {
	t.actions = append(t.actions, action)
}

func (t *timer) AddAlarm(action func(string), when tobject.Time) {
	t.actions = append(t.actions, func(current string) {
		if when.Equal(t.present) {
			action(current)
		}
	})
}

func (t *timer) DoWhenFinished(action func()) {
	t.finalAction = action
}

func (t *timer) Reset() {
	t.present = t.deadline
	t.actions = make([]func(string), 0)
}

func (t *timer) Pause() {
	t.ticker.Stop()
}
