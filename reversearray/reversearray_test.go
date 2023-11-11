package reversearray

import (
	"reflect"
	"sort"
	"testing"
)

var ds = []int{1, 2, 3, 4, 5}

func TestReverseArray(t *testing.T) {
	var (
		e = []int{5, 4, 3, 2, 1}
		r = reverse(ds)
	)

	if !reflect.DeepEqual(e, r) {
		t.Errorf("expected %v, got %v", e, r)
	}
}

func BenchmarkReverseArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(reverse(ds))
	}
}
