package bubblesort

import (
	"reflect"
	"testing"
)

var ds = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}

func TestBubbleSort(t *testing.T) {
	var (
		e   = []int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}
		res = sort(ds[:])
	)

	if !reflect.DeepEqual(e, res) {
		t.Errorf("expected %q, got %q", e, res)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort(ds[:])
	}
}
