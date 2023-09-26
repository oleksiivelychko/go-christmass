package sort

import (
	"reflect"
	"testing"
)

var unsorted = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}

func TestBubbleSort(t *testing.T) {
	sorted := BubbleSort(unsorted[:])
	if !reflect.DeepEqual([]int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}, sorted) {
		t.Error("dataset was not sorted")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(unsorted[:])
	}
}
