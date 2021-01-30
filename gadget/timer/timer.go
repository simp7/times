package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
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
	once        sync.Once
	isRunning   bool
	actions     []func(string)
	finalAction func()
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

func (t *timer) do() {
	current := t.formatter.Format(t.present)
	for _, action := range t.actions {
		action(current)
	}
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

	if t.finalAction != nil {
		t.finalAction()
	}

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
	t.AddAlarm(func(string) { t.Stop() }, tobject.StandardZero())
	t.once = sync.Once{}
}

func (t *timer) Pause() {
	if t.isRunning {
		t.ticker.Stop()
		t.isRunning = false
	}
}
