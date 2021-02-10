package clock

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
	"time"
)

//Clock is an interface that returns current time.
type Clock interface {
	gadget.Gadget
}

type clock struct {
	ticker    gadget.Ticker
	present   tobject.Time
	formatter formatter.TimeFormatter
	unit      tobject.Unit
	once      sync.Once
	isRunning bool
	actions   action.Actions
}

//New returns struct that implements Clock.
//parameter unit is for ticking rate and formatter for formatting time to string.
func New(unit tobject.Unit, formatter formatter.TimeFormatter) Clock {

	c := new(clock)

	c.unit = unit
	c.formatter = formatter
	c.isRunning = false

	c.ticker = gadget.NewTicker(unit)

	c.Reset()

	return c

}

func (c *clock) Start() {
	c.sync()
	c.isRunning = true
	c.work()
}

func (c *clock) getAction() action.Action {
	return c.actions.ActionsWhen(c.present)
}

func (c *clock) do() {
	current := c.Present()
	c.getAction().Do(current)
}

func (c *clock) work() {
	c.ticker.Start(func() {
		c.present.Tick()
		c.do()
	})
}

func (c *clock) Stop() string {

	result := c.Present()

	c.Pause()
	c.Reset()

	return result

}

func (c *clock) Add(f func(current string)) {
	c.actions.Add(action.NewAction(f), nil)
}

func (c *clock) AddAlarm(f func(current string), when tobject.Time) {
	c.actions.Add(action.NewAction(f), when)
}

func (c *clock) Reset() {
	c.actions = action.NewActions()
	c.sync()
}

func (c *clock) sync() {
	if c.unit == tobject.Ms {
		c.present = tobject.AccurateFor(time.Now())
	} else {
		c.present = tobject.StandardFor(time.Now())
	}
}

func (c *clock) Pause() {
	if c.isRunning {
		c.isRunning = false
		c.ticker.Stop()
		c.once = sync.Once{}
	}
}

func (c *clock) Present() string {
	return c.formatter.Format(c.present)
}
