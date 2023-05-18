package reverse_array

import (
	"reflect"
	"sort"
	"testing"
)

func TestArrayReverse(t *testing.T) {
	var slice = make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5

	if !reflect.DeepEqual(ReverseArray(slice), []int{5, 4, 3, 2, 1}) {
		t.Error("arrays are not equal")
	}
}

func BenchmarkReverseArray(b *testing.B) {
	var slice = []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		sort.Ints(ReverseArray(slice))
	}
}
