package stopwatch

import (
	"github.com/simp7/times/gadget"
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
	actions   []func(string)
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

func (s *stopwatch) do() {
	current := s.formatter.Format(s.present)
	for _, action := range s.actions {
		action(current)
	}
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

func (s *stopwatch) Add(action func(string)) {
	s.actions = append(s.actions, action)
}

func (s *stopwatch) AddAlarm(action func(string), when tobject.Time) {
	s.actions = append(s.actions, func(current string) {
		if when.Equal(s.present) {
			action(current)
		}
	})
}

func (s *stopwatch) Reset() {

	s.actions = make([]func(string), 0)

	if s.unit == tobject.Ms {
		s.present = tobject.Accurate(0, 0, 0, 0, 0)
	} else {
		s.present = tobject.Standard(0, 0, 0, 0)
	}

	s.once = sync.Once{}

}

func (s *stopwatch) Pause() {
	if s.isRunning {
		s.isRunning = false
		s.ticker.Stop()
	}
}
