package clock

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
	"time"
)

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

func New(u tobject.Unit, f formatter.TimeFormatter) Clock {

	c := new(clock)

	c.unit = u
	c.formatter = f
	c.isRunning = false

	c.ticker = gadget.NewTicker(u)

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
	current := c.formatter.Format(c.present)
	c.getAction().Do(current)
}

func (c *clock) work() {
	c.ticker.Start(func() {
		c.present.Tick()
		c.do()
	})
}

func (c *clock) Stop() string {

	result := c.formatter.Format(c.present)

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
