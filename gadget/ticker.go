package gadget

import (
	"github.com/simp7/times/model/tobject"
	"time"
)

//Ticker do function in specific period.
type Ticker interface {
	Start(func()) //Start implement function and operate Ticker. When ticker ticks, function of parameter would be called.
	Stop()        //Stop stops Ticker ticking.
}

type ticker struct {
	unit      time.Duration
	stopper   chan struct{}
	isRunning bool
}

//NewTicker Returns Ticker that ticks in rate of unit.
//Parameter unit determines the rate of ticker.
func NewTicker(unit tobject.Unit) Ticker {
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
