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
	unit      time.Duration
	once      sync.Once
	format    times.Format
	isRunning bool
	actions   times.Actions
}

//New returns struct that implements times.Gadget
//parameter unit is for ticking rate and format for formatting time to string.
func New(unit time.Duration, format times.Format) *stopwatch {

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
	s.getAction().Do(s.present)
}

func (s *stopwatch) work() {
	s.ticker.Start(func() {
		s.present.Tick()
		s.once.Do(s.present.Rewind)
		s.do()
	})
}

func (s *stopwatch) Stop() times.Object {

	result := s.present

	s.Pause()
	s.Reset()

	return result

}

func (s *stopwatch) Add(action times.Action) {
	s.actions.Add(action, nil)
}

func (s *stopwatch) AddAlarm(action times.Action, when times.Object) {
	s.actions.Add(action, when)
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

func (s *stopwatch) Format(obj times.Object) string {
	return s.format(obj)
}
