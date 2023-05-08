package sort

import (
	"github.com/oleksiivelychko/go-code-helpers/array"
	"testing"
)

var unsortedIntArr = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}

func TestBubbleSort(t *testing.T) {
	sorted := BubbleSort(unsortedIntArr[:])
	if !array.Equal(unsortedIntArr[:], sorted) {
		t.Error("arrays are not equal")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(unsortedIntArr[:])
	}
}
