package main

import (
	"fmt"
	"github.com/simp7/times"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/gadget/clock"
	"github.com/simp7/times/object/formatter"
)

func main() {
	c := clock.Standard
	a := action.NewAction(func(object times.Object) {
		fmt.Println(formatter.Clock.Notation12(object))
	})
	c.Add(a)
	c.Start()
}
