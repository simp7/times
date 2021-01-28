package stopwatch

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
	"time"
)

//Stopwatch is an interface that set deadline and runs until deadline has been passed or End is called.
type Stopwatch interface {
	gadget.Gadget
}

type stopwatch struct {
	ticker    time.Ticker
	present   tobject.Time
	formatter formatter.TimeFormatter
	stopper   chan struct{}
	unit      tobject.Unit
	actions   []func(string)
	once      sync.Once
}

func New(u tobject.Unit, f formatter.TimeFormatter) Stopwatch {

	s := new(stopwatch)

	s.unit = u
	s.formatter = f
	s.actions = make([]func(string), 0)

	return s

}

func (s *stopwatch) Start() {

	s.ticker = *time.NewTicker(time.Duration(s.unit))
	s.stopper = make(chan struct{})

	if s.unit == tobject.Ms {
		s.present = tobject.Accurate(0, 0, 0, 0, 0)
	} else {
		s.present = tobject.Standard(0, 0, 0, 0)
	}
	s.once.Do(func() {
		s.do()
		go s.working()
	})
	<-s.stopper

}

func (s *stopwatch) do() {
	current := s.formatter.Format(s.present)
	for _, action := range s.actions {
		action(current)
	}
}

func (s *stopwatch) working() {

	for {
		select {

		case <-s.ticker.C:
			s.present.Tick()
			s.do()

		case <-s.stopper:
			s.ticker.Stop()
			return

		}
	}

}

func (s *stopwatch) End() string {

	close(s.stopper)
	return s.formatter.Format(s.present)

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

func (s *stopwatch) Reset() Stopwatch {
	return New(s.unit, s.formatter)
}
