package stopwatch

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/object"
	"sync"
	"time"
)

type stopwatch struct {
	ticker    *gadget.Ticker
	present   times.Object
	format    times.Format
	unit      time.Duration
	once      sync.Once
	isRunning bool
	actions   times.Actions
}

//New returns struct that implements times.Gadget
//parameter unit is for ticking rate and format for formatting time to string.
func New(unit time.Duration, format times.Format) times.Gadget {

	s := new(stopwatch)

	s.unit = unit
	s.format = format
	s.isRunning = false

	s.ticker = gadget.NewTicker(unit)

	s.Reset()

	return s

}

func (s *stopwatch) Start() {
	s.isRunning = true
	s.work()
}

func (s *stopwatch) getAction() times.Action {
	return s.actions.ActionsWhen(s.present)
}

func (s *stopwatch) do() {
	current := s.Present()
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

	result := s.Present()

	s.Pause()
	s.Reset()

	return result

}

func (s *stopwatch) Add(f func(string)) {
	s.actions.Add(action.NewAction(f), nil)
}

func (s *stopwatch) AddAlarm(f func(string), when times.Object) {
	s.actions.Add(action.NewAction(f), when)
}

func (s *stopwatch) Reset() {

	s.actions = action.NewActions()

	if s.unit == time.Millisecond {
		s.present = object.Accurate(0, 0, 0, 0, 0)
	} else {
		s.present = object.Standard(0, 0, 0, 0)
	}

}

func (s *stopwatch) Pause() {
	if s.isRunning {
		s.isRunning = false
		s.ticker.Stop()
		s.once = sync.Once{}
	}
}

func (s *stopwatch) Present() string {
	return s.format(s.present)
}
