package timer

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
	"time"
)

//Timer is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Timer interface {
	gadget.Gadget
	DoWhenFinished(func()) //DoWhenFinished is called when time of timer becomes zero.
}

type timer struct {
	ticker      time.Ticker
	present     tobject.Time
	formatter   formatter.TimeFormatter
	stopper     chan struct{}
	unit        tobject.Unit
	actions     []func(string)
	finalAction func()
	once        sync.Once
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

	t.once.Do(func() {
		t.do()
		go t.working()
	})
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
				t.Stop()
			}

		case <-t.stopper:
			t.ticker.Stop()
			return

		}
	}

}

func (t *timer) Stop() string {

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

func (t *timer) Reset() {
	//TODO: Implement me.
}

func (t *timer) Pause() {
	//TODO: Implement me.
}
