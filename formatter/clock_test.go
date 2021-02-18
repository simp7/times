package formatter

import "testing"

func TestClockFormatter_Format(t *testing.T) {
	testSkeleton(Clock(false), standardPreset(), []string{"00:00:00", "00:00:05", "00:10:03", "03:02:01", "03:02:01", "00:09:10", "18:09:00", "12:30:00"}, t)
	testSkeleton(Clock(true), standardPreset(), []string{"12:00:00 AM", "12:00:05 AM", "12:10:03 AM", "3:02:01 AM", "3:02:01 AM", "12:09:10 AM", "6:09:00 PM", "12:30:00 PM"}, t)
}
