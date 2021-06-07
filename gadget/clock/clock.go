package clock

import (
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/object"
	"sync"
	"time"
)

type clock struct {
	ticker    *gadget.Ticker
	present   times.Object
	format    times.Format
	unit      time.Duration
	once      sync.Once
	isRunning bool
	actions   times.Actions
}

//New returns struct that implements times.Gadget.
//parameter unit is for ticking rate and format for formatting time to string.
func New(unit time.Duration, format times.Format) *clock {

	c := new(clock)

	c.unit = unit
	c.format = format
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

func (c *clock) getAction() times.Action {
	return c.actions.ActionsWhen(c.present)
}

func (c *clock) do() {
	c.getAction().Do(c.present)
}

func (c *clock) work() {
	c.ticker.Start(func() {
		c.present.Tick()
		c.do()
	})
}

func (c *clock) Stop() times.Object {

	result := c.present

	c.Pause()
	c.Reset()

	return result

}

func (c *clock) Add(action times.Action) {
	c.actions.Add(action, nil)
}

func (c *clock) AddAlarm(action times.Action, when times.Object) {
	c.actions.Add(action, when)
}

func (c *clock) Reset() {
	c.actions = action.NewActions()
	c.sync()
}

func (c *clock) sync() {
	if c.unit == time.Millisecond {
		c.present = object.AccurateFor(time.Now())
	} else {
		c.present = object.StandardFor(time.Now())
	}
}

func (c *clock) Pause() {
	if c.isRunning {
		c.isRunning = false
		c.ticker.Stop()
		c.once = sync.Once{}
	}
}

func (c *clock) Format(obj times.Object) string {
	return c.format(obj)
}
