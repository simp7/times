package main

import (
	"fmt"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/model/tobject"
	"github.com/simp7/times/sample/preset"
	"os"
	"strconv"
)

func recursiveCall(gadgets []gadget.Gadget, a1, a2 tobject.Time) {

	if len(gadgets) == 0 {
		return
	}

	this := gadgets[0]

	this.Add(func(current string) {
		fmt.Println(current)
	})

	this.AddAlarm(func(string) {
		fmt.Println("Paused")
		this.Pause()
		recursiveCall(gadgets[1:], a1, a2)
		this.Start()
	}, a1)

	this.AddAlarm(func(string) {
		fmt.Println("Finished!")
		this.Stop()
	}, a2)

	this.Start()

}

func changeStopwatchesToGadgets(stopwatches []stopwatch.Stopwatch) []gadget.Gadget {
	gadgets := make([]gadget.Gadget, len(stopwatches))
	for i := range stopwatches {
		gadgets[i] = stopwatches[i]
	}
	return gadgets
}

func changeTimersToGadgets(timers []timer.Timer) []gadget.Gadget {
	gadgets := make([]gadget.Gadget, len(timers))
	for i := range timers {
		gadgets[i] = timers[i]
	}
	return gadgets
}

func main() {

	second := 5 //default value

	if len(os.Args) == 2 {
		var err error
		second, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argument is not number.")
			os.Exit(3)
		}
		if second > 30 {
			fmt.Println("Second should less than 30")
			second = 29
		}
	}

	alarm1 := tobject.Standard(second, 0, 0, 0)
	alarm2 := tobject.Standard(second*2, 0, 0, 0)

	recursiveCall(changeStopwatchesToGadgets(preset.Stopwatches()), alarm1, alarm2)
	recursiveCall(changeTimersToGadgets(preset.Timers(second*2)), alarm1, tobject.StandardZero().SetSecond(1))

}
