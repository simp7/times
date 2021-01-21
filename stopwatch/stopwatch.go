package stopwatch

import (
	"time"
	"times/model/formatter"
	"times/model/gadget"
	"times/model/tobject"
)

type Stopwatch interface {
	gadget.Gadget
}

type stopwatch struct {
	time.Ticker
	present   tobject.Time
	formatter formatter.TimeFormatter
	stopper   chan struct{}
	unit      tobject.Unit
	actions   []func(string)
}

func New(u tobject.Unit, f formatter.TimeFormatter) Stopwatch {

	s := new(stopwatch)

	s.stopper = make(chan struct{})
	s.unit = u
	s.formatter = f
	s.actions = make([]func(string), 0)

	return s

}

func (s *stopwatch) Start() {

	s.Ticker = *time.NewTicker(time.Duration(s.unit))

	if s.unit == tobject.Ms {
		s.present = tobject.Accurate(0,0,0,0,0)
	} else {
		s.present = tobject.Standard(0,0,0,0)
	}

	go s.working()
	<- s.stopper

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

		case <-s.C:
			s.present.Tick()
			s.do()

		case <-s.stopper:
			s.Stop()
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
	s.actions = append(s.actions, func(current string){
		if when.Equal(s.present) {
			action(current)
		}
	})
}