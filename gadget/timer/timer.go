package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"time"
)

type Timer interface {
	gadget.Gadget
	DoWhenFinished(func())
}

type timer struct {
	ticker      time.Ticker
	present     tobject.Time
	formatter   formatter.TimeFormatter
	stopper     chan struct{}
	unit        tobject.Unit
	actions     []func(string)
	finalAction func()
}

func New(u tobject.Unit, f formatter.TimeFormatter, deadline tobject.Time) Timer {

	t := new(timer)

	t.unit = u
	t.formatter = f
	t.present = deadline
	t.actions = make([]func(string), 0)

	return t

}

func (t *timer) Start() {

	t.ticker = *time.NewTicker(time.Duration(t.unit))
	t.stopper = make(chan struct{})

	t.do()

	go t.working()
	<-t.stopper

}

func (t *timer) do() {
	current := t.formatter.Format(t.present)
	for _, action := range t.actions {
		action(current)
	}
}

func (t *timer) working() {

	for {
		select {

		case <-t.ticker.C:

			t.present.Rewind()
			t.do()

			if t.present.Equal(tobject.AccurateZero()) {

				if t.finalAction != nil {
					t.finalAction()
				}

				t.End()

			}

		case <-t.stopper:
			t.ticker.Stop()
			return

		}
	}

}

func (t *timer) End() string {

	close(t.stopper)
	return t.formatter.Format(t.present)

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
