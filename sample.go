package main

import (
	"fmt"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"github.com/simp7/times/gadget/timer"
	"github.com/simp7/times/model/tobject"
)

func getStandard() tobject.Time {
	return tobject.StandardZero().SetSecond(5)
}

func getAccurate() tobject.Time {
	return tobject.AccurateZero().SetSecond(5)
}

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
			s.Stop()
		}, getStandard())
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

	for _, s := range []timer.Timer{timer.Standard(getStandard()), timer.Hangul(getStandard()), timer.Detail(getAccurate())} {
		testTimer(s)
	}

}
