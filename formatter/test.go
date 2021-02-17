package formatter

import (
	"errors"
	"github.com/simp7/times"
	"github.com/simp7/times/ttime"
	"testing"
)

func standardPreset() []times.Time {
	return []times.Time{ttime.StandardZero(), ttime.Standard(5, 0, 0, 0), ttime.Standard(3, 10, 0, 0), ttime.Standard(1, 2, 3, 0), ttime.Standard(1, 2, 3, 4), ttime.Standard(10, 9, 0, 8), ttime.Standard(0, 9, 18, 27)}
}

func accuratePreset() []times.Time {
	return []times.Time{ttime.AccurateZero(), ttime.Accurate(0, 5, 0, 0, 0), ttime.Accurate(0, 3, 10, 0, 0), ttime.Accurate(0, 1, 2, 3, 0), ttime.Accurate(0, 1, 2, 3, 4), ttime.Accurate(0, 10, 9, 0, 8), ttime.Accurate(7, 0, 0, 0, 0), ttime.Accurate(42, 50, 3, 11, 0), ttime.Accurate(999, 0, 0, 0, 1)}
}

func compare(idx int, want, get string, t *testing.T) {
	if want == get {
		t.Logf("Successfully passed in example %d\n", idx)
	} else {
		t.Errorf("Errors in Example %d : wanted %s, get %s\n", idx, want, get)
	}
}

func testSkeleton(f times.TimeFormatter, preset []times.Time, answer []string, t *testing.T) {
	if len(answer) != len(preset) {
		panic(differentLen)
	}
	for i, v := range preset {
		compare(i, answer[i], f.Format(v), t)
	}
}

var (
	differentLen = errors.New("the length of array between test case and answer is different")
)
