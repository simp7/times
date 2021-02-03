package clock

import (
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/model/action"
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
	"sync"
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
	c.Start()

	return c

}

func (c *clock) Start() {
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

	if c.unit == tobject.Ms {
		c.present = tobject.Accurate(0, 0, 0, 0, 0)
	} else {
		c.present = tobject.Standard(0, 0, 0, 0)
	}

	c.sync()

}

func (c *clock) Pause() {
	if c.isRunning {
		c.isRunning = false
		c.ticker.Stop()
		c.once = sync.Once{}
	}
}

func (c *clock) sync() {
	//TODO: synchronize clock to now.
}
