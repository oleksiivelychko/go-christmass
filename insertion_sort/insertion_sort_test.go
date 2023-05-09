package insertion_sort

import (
	"github.com/oleksiivelychko/go-code-helpers/array"
	"testing"
)

var unsorted = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}

func TestInsertionSort(t *testing.T) {
	var sorted = InsertionSort(unsorted[:])
	if !array.Equal(unsorted[:], sorted) {
		t.Error("arrays are not equal")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(unsorted[:])
	}
}
