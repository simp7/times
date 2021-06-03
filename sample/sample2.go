package main

import (
	"fmt"
	"github.com/simp7/times"
	"github.com/simp7/times/gadget/action"
	"github.com/simp7/times/object"
	"github.com/simp7/times/object/formatter"
	"github.com/simp7/times/sample/preset"
	"os"
	"strconv"
)

func recursiveCall(gadgets []times.Gadget, obj1, obj2 times.Object) {

	if len(gadgets) == 0 {
		return
	}

	this := gadgets[0]

	action1 := action.NewAction(func(o times.Object) {
		fmt.Println(formatter.Detail(o))
	})
	action2 := action.NewAction(func(times.Object) {
		fmt.Println("Paused")
		this.Pause()
		recursiveCall(gadgets[1:], obj1, obj2)
		this.Start()
	})
	action3 := action.NewAction(func(times.Object) {
		fmt.Println("Finished!")
		this.Stop()
	})

	this.Add(action1)
	this.AddAlarm(action2, obj1)
	this.AddAlarm(action3, obj2)

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

	alarm1 := object.Standard(second, 0, 0, 0)
	alarm2 := object.Standard(second*2, 0, 0, 0)

	recursiveCall(changeStopwatchesToGadgets(preset.Stopwatches()), alarm1, alarm2)
	recursiveCall(changeTimersToGadgets(preset.Timers(second*2)), alarm1, object.StandardZero().SetSecond(0))

}
