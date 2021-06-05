package main

import (
	"fmt"
	"github.com/simp7/times"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/gadget/clock"
)

func main() {
	c := clock.Standard
	a := action.NewAction(func(object times.Object) {
		fmt.Println(c.Format(object))
	})
	c.Add(a)
	c.Start()
}
