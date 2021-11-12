package smallest_search

import (
	"testing"
)

var slice = []int{10, 1, -1, -2, 2}

func TestSmallestSearch(t *testing.T) {
	result := SmallestSearch(slice, slice[0])
	if result != -2 {
		t.Errorf("[func SmallestSearch(slice []int) int] -> %d != -2", result)
	}
}

func BenchmarkSmallestSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallestSearch(slice, slice[0])
	}
}
