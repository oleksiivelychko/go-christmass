package sort

import (
	"reflect"
	"sort"
	"testing"
)

var arrInt = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}
var arrString = [10]string{"f", "h", "a", "abc", "x", "o", "1", "u", "12", "s"}

func TestQuickSort(t *testing.T) {
	var sorted = QuickSort(arrInt[:])
	if !reflect.DeepEqual([]int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}, sorted) {
		t.Error("arrays are equal")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(arrInt[:])
	}
}

func TestQuickSort_BuiltinInt(t *testing.T) {
	sort.Ints(arrInt[:])

	if !sort.IntsAreSorted(arrInt[:]) {
		t.Error("array is not sorted")
	}

	t.Log(arrInt)
}

func BenchmarkQuickSort_BuiltinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(arrInt[:])
	}
}

func TestQuickSort_BuiltinString(t *testing.T) {
	sort.Strings(arrString[:])

	if !sort.StringsAreSorted(arrString[:]) {
		t.Error("array is not sorted")
	}

	t.Log(arrString)
}

func BenchmarkQuickSort_BuiltinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Strings(arrString[:])
	}
}
