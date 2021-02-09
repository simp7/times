package main

import (
	"fmt"
	"github.com/simp7/times/gadget/clock"
)

func main() {
	c := clock.Standard

	c.Add(func(s string) {
		fmt.Println(s)
	})

	c.Start()
}
