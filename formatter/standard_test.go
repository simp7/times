package formatter

import (
	"testing"
)

func TestStandardFormatter_Format(t *testing.T) {
	testSkeleton(Standard(), standardPreset(), []string{"00:00", "00:05", "10:03", "03:02:01", "4:03:02:01", "8:00:09:10", "27:18:09:00"}, t)
}
