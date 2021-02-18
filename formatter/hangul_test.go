package formatter

import "testing"

func TestHangulFormatter_Format(t *testing.T) {
	testSkeleton(Hangul(), standardPreset(), []string{"0초", "5초", "10분 3초", "3시간 2분 1초", "4일 3시간 2분 1초", "8일 9분 10초", "27일 18시간 9분", "4일 12시간 30분"}, t)
}
