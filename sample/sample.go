package main

import (
	"fmt"
	"github.com/simp7/times"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/sample/preset"
	"github.com/simp7/times/time"
	"os"
	"strconv"
)

func test1Skeleton(g times.Gadget, someFunc func()) {
	g.Add(func(current string) {
		fmt.Println(current)
	})
	someFunc()
	g.Start()
}

func testStopwatch1(s gadget.Stopwatch, t int) {
	test1Skeleton(s, func() {
		s.AddAlarm(func(current string) {
			fmt.Println("Finished!")
			s.Stop()
		}, preset.GetStandard(t))
	})
}

func testTimer1(t gadget.Timer) {
	test1Skeleton(t, func() {
		t.AddAlarm(func(string) {
			fmt.Println("Finished!")
		}, time.StandardZero())
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
