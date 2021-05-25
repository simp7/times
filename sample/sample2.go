package main

import (
	"fmt"
	"github.com/simp7/times"
	"github.com/simp7/times/sample/preset"
	"github.com/simp7/times/timeObject"
	"os"
	"strconv"
)

func recursiveCall(gadgets []times.Gadget, a1, a2 times.Time) {

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

func changeStopwatchesToGadgets(stopwatches []times.Gadget) []times.Gadget {
	gadgets := make([]times.Gadget, len(stopwatches))
	for i := range stopwatches {
		gadgets[i] = stopwatches[i]
	}
	return gadgets
}

func changeTimersToGadgets(timers []times.Gadget) []times.Gadget {
	gadgets := make([]times.Gadget, len(timers))
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

	alarm1 := timeObject.Standard(second, 0, 0, 0)
	alarm2 := timeObject.Standard(second*2, 0, 0, 0)

	recursiveCall(changeStopwatchesToGadgets(preset.Stopwatches()), alarm1, alarm2)
	recursiveCall(changeTimersToGadgets(preset.Timers(second*2)), alarm1, timeObject.StandardZero().SetSecond(0))

}
