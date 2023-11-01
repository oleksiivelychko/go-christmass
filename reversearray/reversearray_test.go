package reversearray

import (
	"reflect"
	"sort"
	"testing"
)

func TestReverseArray(t *testing.T) {
	var ds = make([]int, 5)
	ds[0] = 1
	ds[1] = 2
	ds[2] = 3
	ds[3] = 4
	ds[4] = 5

	if !reflect.DeepEqual(reverse(ds), []int{5, 4, 3, 2, 1}) {
		t.Error("unable to reverse")
	}
}

func BenchmarkReverseArray(b *testing.B) {
	var ds = []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		sort.Ints(reverse(ds))
	}
}
