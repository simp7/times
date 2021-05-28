package gadget

import (
	"time"
)

//Ticker is an struct that ticks after fixed duration.
type Ticker struct {
	unit      time.Duration
	stopper   chan struct{}
	isRunning bool
}

//NewTicker returns Ticker that ticks in rate of unit and parameter determines the rate of Ticker.
func NewTicker(unit time.Duration) *Ticker {
	t := new(Ticker)
	t.unit = unit
	t.isRunning = false
	return t
}

//Start makes ticker ticks.
func (t *Ticker) Start(action func()) {

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

//Stop stops ticker to tick.
func (t *Ticker) Stop() {
	if t.isRunning {
		t.isRunning = false
		close(t.stopper)
	}
}
