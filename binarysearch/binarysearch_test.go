package binarysearch

import (
	"math/rand"
	"sort"
	"testing"
)

var intDS = [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
var floatDS = [10]float64{1, 1.5, 3, 3.5, 5, 5.5, 7, 7.5, 9, 9.5}
var stringDS = [10]string{"1", "12", "a", "abc", "f", "h", "o", "s", "u", "x"}

func TestBinarySearch_Int(t *testing.T) {
	found, index := search(intDS[:], 5)
	if !found || index != 2 {
		t.Errorf("unable to find, got index %d", index)
	}
}

func BenchmarkBinarySearch_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_Float(t *testing.T) {
	found, index := search(floatDS[:], 7.5)
	if !found || index != 7 {
		t.Errorf("unable to find, got index %d", index)
	}
}

func BenchmarkBinarySearch_Float(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_RecursionInt(t *testing.T) {
	found := searchRecursion(intDS[:], 5)
	if !found {
		t.Error("unable to find")
	}
}

func BenchmarkBinarySearch_RecursionInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRecursion(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_RecursionFloat(t *testing.T) {
	found := searchRecursion(floatDS[:], 7.5)
	if !found {
		t.Error("unable to find")
	}
}

func BenchmarkBinarySearch_RecursionFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRecursion(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinInt(t *testing.T) {
	index := sort.SearchInts(intDS[:], 5)
	if index != 2 {
		t.Errorf("unable to find, got index %d", index)
	}
}

func BenchmarkBinarySearch_BuiltinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_BuiltinFloat(t *testing.T) {
	index := sort.SearchFloat64s(floatDS[:], 7.5)
	if index != 7 {
		t.Errorf("unable to find, got index %d", index)
	}
}

func BenchmarkBinarySearch_BuiltinFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchFloat64s(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinString(t *testing.T) {
	index := sort.SearchStrings(stringDS[:], "abc")
	if index != 3 {
		t.Errorf("got invalid index %d", index)
	}
}

func BenchmarkBinarySearch_BuiltinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchStrings(stringDS[:], "abc")
	}
}
