package insertion_sort

import (
	"reflect"
	"testing"
)

var unsorted = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}

func TestInsertionSort(t *testing.T) {
	var sorted = insertionSort(unsorted[:])
	if !reflect.DeepEqual([]int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}, sorted) {
		t.Error("dataset was not sorted")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort(unsorted[:])
	}
}
