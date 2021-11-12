package quick_sort

import (
	"testing"
)

var array = [3]int{21, -1, 11}

func TestQuickSort(t *testing.T) {
	var sorted = QuickSort(array[:])
	if sorted[0] != -1 {
		t.Errorf("[func QuickSort(slice []int) []int] -> %d != -1", sorted[0])
	}
	if sorted[1] != 11 {
		t.Errorf("[func QuickSort(slice []int) []int] -> %d != 11", sorted[1])
	}
	if sorted[2] != 21 {
		t.Errorf("[func QuickSort(slice []int) []int] -> %d != 21", sorted[2])
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(array[:])
	}
}
