package formatter

import (
	"testing"
)

func TestDetailFormatter_Format(t *testing.T) {
	testSkeleton(Detail(), accuratePreset(), []string{"00:00:000", "00:05:000", "10:03:000", "03:02:01:000", "4:03:02:01:000", "8:00:09:10:000", "00:00:007", "11:03:50:042", "1:00:00:00:999"}, t)
}
