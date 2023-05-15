package binary_search

import (
	"math/rand"
	"sort"
	"testing"
)

var arrInt = [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
var arrFloat = [10]float64{1, 1.5, 3, 3.5, 5, 5.5, 7, 7.5, 9, 9.5}
var arrString = [10]string{"1", "12", "a", "abc", "f", "h", "o", "s", "u", "x"}

func TestBinarySearch_Int(t *testing.T) {
	found, index := BinarySearch(arrInt[:], 5)
	if !found || index != 2 {
		t.Errorf("unable to find item, index is %d", index)
	}
}

func BenchmarkBinarySearch_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(arrInt[:], rand.Intn(7))
	}
}

func TestBinarySearch_Float(t *testing.T) {
	found, index := BinarySearch(arrFloat[:], 7.5)
	if !found || index != 7 {
		t.Errorf("unable to find item, index is %d", index)
	}
}

func BenchmarkBinarySearch_Float(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(arrFloat[:], rand.Float64())
	}
}

func TestBinarySearch_RecursionInt(t *testing.T) {
	found := BinarySearchRecursion(arrInt[:], 5)
	if !found {
		t.Error("unable to find item")
	}
}

func BenchmarkBinarySearch_RecursionInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchRecursion(arrInt[:], rand.Intn(7))
	}
}

func TestBinarySearch_RecursionFloat(t *testing.T) {
	found := BinarySearchRecursion(arrFloat[:], 7.5)
	if !found {
		t.Error("unable to find item")
	}
}

func BenchmarkBinarySearch_RecursionFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchRecursion(arrFloat[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinInt(t *testing.T) {
	index := sort.SearchInts(arrInt[:], 5)
	if index != 2 {
		t.Errorf("unable to find item, index is %d", index)
	}
}

func BenchmarkBinarySearch_BuiltinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(arrInt[:], rand.Intn(7))
	}
}

func TestBinarySearch_BuiltinFloat(t *testing.T) {
	index := sort.SearchFloat64s(arrFloat[:], 7.5)
	if index != 7 {
		t.Errorf("unable to find item, index is %d", index)
	}
}

func BenchmarkBinarySearch_BuiltinFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchFloat64s(arrFloat[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinString(t *testing.T) {
	index := sort.SearchStrings(arrString[:], "abc")
	if index != 3 {
		t.Errorf("index %d is wrong", index)
	}
}

func BenchmarkBinarySearch_BuiltinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchStrings(arrString[:], "abc")
	}
}
