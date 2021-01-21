package main

import (
	"fmt"
	"times/model/gadget"
	"times/model/tobject"
	"times/stopwatch"
	"times/timer"
)

func testSkeleton(g gadget.Gadget, someFunc func()) {
	g.Add(func(current string) {
		fmt.Println(current)
	})
	someFunc()
	g.Start()
}

func testStopwatch(s stopwatch.Stopwatch) {
	testSkeleton(s, func() {
		s.AddAlarm(func(current string) {
			fmt.Println("Finished!")
			s.End()
		}, tobject.Standard(5,0,0,0))
	})
}

func testTimer(t timer.Timer) {
	testSkeleton(t, func() {
		t.DoWhenFinished(func() {
			fmt.Println("Finished!")
		})
	})
}

func main() {

	for _, s := range []stopwatch.Stopwatch{stopwatch.Standard, stopwatch.Hangul, stopwatch.Detail} {
		testStopwatch(s)
	}

	for _, s := range []timer.Timer{timer.Standard(tobject.Standard(5,0,0,0)), timer.Hangul(tobject.Standard(5, 0, 0, 0)), timer.Detail(tobject.Accurate(0, 5, 0,0, 0))} {
		testTimer(s)
	}

}
