package reversed_array

import (
	"reflect"
	"sort"
	"testing"
)

func TestReversedArray(t *testing.T) {
	var dataset = make([]int, 5)
	dataset[0] = 1
	dataset[1] = 2
	dataset[2] = 3
	dataset[3] = 4
	dataset[4] = 5

	if !reflect.DeepEqual(reversedArray(dataset), []int{5, 4, 3, 2, 1}) {
		t.Error("array was not reversed")
	}
}

func BenchmarkReversedArray(b *testing.B) {
	var dataset = []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		sort.Ints(reversedArray(dataset))
	}
}
