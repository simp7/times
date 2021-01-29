package gadget

import (
	"github.com/simp7/times/model/tobject"
	"time"
)

type Ticker interface {
	Start(func())
	Stop()
}

func NewTicker(unit tobject.Unit) Ticker {
	t := new(ticker)
	t.unit = time.Duration(unit)
	return t
}

type ticker struct {
	unit    time.Duration
	stopper chan struct{}
}

func (t *ticker) Start(action func()) {

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

func (t *ticker) Stop() {
	close(t.stopper)
}
