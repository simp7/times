package formatter

import "testing"

func TestClockFormatter_Format(t *testing.T) {
	testSkeleton(Clock(false), standardPreset(), []string{"00:00:00", "00:00:05", "00:10:03", "03:02:01", "03:02:01", "00:09:10", "18:09:00"}, t)
	testSkeleton(Clock(true), standardPreset(), []string{"0:00:00 AM", "0:00:05 AM", "0:10:03 AM", "3:02:01 AM", "3:02:01 AM", "0:09:10 AM", "6:09:00 PM"}, t)
}
