package stopwatch

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
)

//Stopwatch is an interface that set deadline and runs until deadline has been passed or Stop is called.
type Stopwatch interface {
	gadget.Gadget
}

type stopwatch struct {
	ticker    gadget.Ticker
	present   tobject.Time
	formatter formatter.TimeFormatter
	unit      tobject.Unit
	once      sync.Once
	isRunning bool
	actions   action.Actions
}

func New(u tobject.Unit, f formatter.TimeFormatter) Stopwatch {

	s := new(stopwatch)

	s.unit = u
	s.formatter = f
	s.isRunning = false

	s.ticker = gadget.NewTicker(u)

	s.Reset()

	return s

}

func (s *stopwatch) Start() {
	s.isRunning = true
	s.work()
}

func (s *stopwatch) getAction() action.Action {
	return s.actions.ActionsWhen(s.present)
}

func (s *stopwatch) do() {
	current := s.formatter.Format(s.present)
	s.getAction().Do(current)
}

func (s *stopwatch) work() {
	s.ticker.Start(func() {
		s.present.Tick()
		s.once.Do(s.present.Rewind)
		s.do()
	})
}

func (s *stopwatch) Stop() string {

	result := s.formatter.Format(s.present)

	s.Pause()
	s.Reset()

	return result

}

func (s *stopwatch) Add(f func(string)) {
	s.actions.Add(action.NewAction(f), nil)
}

func (s *stopwatch) AddAlarm(f func(string), when tobject.Time) {
	s.actions.Add(action.NewAction(f), when)
}

func (s *stopwatch) Reset() {

	s.actions = action.NewActions()

	if s.unit == tobject.Ms {
		s.present = tobject.Accurate(0, 0, 0, 0, 0)
	} else {
		s.present = tobject.Standard(0, 0, 0, 0)
	}

}

func (s *stopwatch) Pause() {
	if s.isRunning {
		s.isRunning = false
		s.ticker.Stop()
		s.once = sync.Once{}
	}
}
