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

func test1Skeleton(g gadget.Gadget, someFunc func()) {

	g.Add(func(current string) {
		fmt.Println(current)
	})
	someFunc()
	g.Start()

}

func testStopwatch1(s stopwatch.Stopwatch, t int) {

	test1Skeleton(s, func() {
		s.AddAlarm(func(current string) {
			fmt.Println("Finished!")
			s.Stop()
		}, preset.GetStandard(t))
	})

}

func testTimer1(t timer.Timer) {

	test1Skeleton(t, func() {
		t.AddAlarm(func(string) {
			fmt.Println("Finished!")
		}, tobject.StandardZero())
	})

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
	}

	for _, s := range preset.Stopwatches() {
		testStopwatch1(s, second)
	}

	for _, t := range preset.Timers(second) {
		testTimer1(t)
	}

}
