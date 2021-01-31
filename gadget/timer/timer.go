package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
)

//Timer is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Timer interface {
	gadget.Gadget
}

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

func New(u tobject.Unit, f formatter.TimeFormatter, deadline tobject.Time) Timer {

	t := new(timer)

	t.unit = u
	t.formatter = f
	t.deadline = deadline
	t.isRunning = false

	t.ticker = gadget.NewTicker(u)

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
	current := t.formatter.Format(t.present)
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

	result := t.formatter.Format(t.present)

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
	t.present = t.deadline
	t.actions = action.NewActions()
	t.AddAlarm(func(string) { t.Stop() }, tobject.StandardZero())
}

func (t *timer) Pause() {
	if t.isRunning {
		t.ticker.Stop()
		t.isRunning = false
		t.once = sync.Once{}
	}
}
