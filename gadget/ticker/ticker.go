package ticker

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"time"
)

type ticker struct {
	unit      time.Duration
	stopper   chan struct{}
	isRunning bool
}

//NewTicker returns struct that implements gadget.Ticker.
//Returned ticker ticks in rate of unit and parameter determines the rate of ticker.
func NewTicker(unit times.Unit) gadget.Ticker {
	t := new(ticker)
	t.unit = time.Duration(unit)
	t.isRunning = false
	return t
}

func (t *ticker) Start(action func()) {

	if !t.isRunning {

		t.isRunning = true
		t.stopper = make(chan struct{})

		action()

		for {
			select {
			case <-time.After(t.unit):
				action()
			case <-t.stopper:
				return
			}
		}

	}

}

func (t *ticker) Stop() {
	if t.isRunning {
		t.isRunning = false
		close(t.stopper)
	}
}
