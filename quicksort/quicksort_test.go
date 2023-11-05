package quicksort

import (
	"reflect"
	goSort "sort"
	"testing"
)

var (
	intDS    = [10]int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}
	stringDS = [10]string{"f", "h", "a", "abc", "x", "o", "1", "u", "12", "s"}
)

func TestQuickSort(t *testing.T) {
	var (
		e      = []int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}
		sorted = sort(intDS[:])
	)

	if !reflect.DeepEqual(e, sorted) {
		t.Errorf("expected %v, got %v", e, sorted)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort(intDS[:])
	}
}

func TestQuickSort_BuiltinInt(t *testing.T) {
	goSort.Ints(intDS[:])

	if !goSort.IntsAreSorted(intDS[:]) {
		t.Error("unable to sort")
	}

	t.Log(intDS)
}

func BenchmarkQuickSort_BuiltinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goSort.Ints(intDS[:])
	}
}

func TestQuickSort_BuiltinString(t *testing.T) {
	goSort.Strings(stringDS[:])

	if !goSort.StringsAreSorted(stringDS[:]) {
		t.Error("unable to sort")
	}

	t.Log(stringDS)
}

func BenchmarkQuickSort_BuiltinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goSort.Strings(stringDS[:])
	}
}
