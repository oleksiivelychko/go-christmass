package binarysearch

import (
	"math/rand"
	"sort"
	"testing"
)

var (
	intDS    = [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	floatDS  = [10]float64{1, 1.5, 3, 3.5, 5, 5.5, 7, 7.5, 9, 9.5}
	stringDS = [10]string{"1", "12", "a", "abc", "f", "h", "o", "s", "u", "x"}
)

func TestBinarySearch_Int(t *testing.T) {
	if ok, i := search(intDS[:], 5); !ok || i != 2 {
		t.Errorf("expected %d, got %d", 2, i)
	}
}

func BenchmarkBinarySearch_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_Float(t *testing.T) {
	if ok, i := search(floatDS[:], 7.5); !ok || i != 7 {
		t.Errorf("expected %d, got %d", 7, i)
	}
}

func BenchmarkBinarySearch_Float(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_RecursionInt(t *testing.T) {
	if !searchRecursion(intDS[:], 5) {
		t.Error("unable to find")
	}
}

func BenchmarkBinarySearch_RecursionInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRecursion(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_RecursionFloat(t *testing.T) {
	if !searchRecursion(floatDS[:], 7.5) {
		t.Error("unable to find")
	}
}

func BenchmarkBinarySearch_RecursionFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRecursion(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinInt(t *testing.T) {
	if i := sort.SearchInts(intDS[:], 5); i != 2 {
		t.Errorf("expected %d, got %d", 2, i)
	}
}

func BenchmarkBinarySearch_BuiltinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(intDS[:], rand.Intn(7))
	}
}

func TestBinarySearch_BuiltinFloat(t *testing.T) {
	if i := sort.SearchFloat64s(floatDS[:], 7.5); i != 7 {
		t.Errorf("expected %d, got %d", 7, i)
	}
}

func BenchmarkBinarySearch_BuiltinFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchFloat64s(floatDS[:], rand.Float64())
	}
}

func TestBinarySearch_BuiltinString(t *testing.T) {
	if i := sort.SearchStrings(stringDS[:], "abc"); i != 3 {
		t.Errorf("expected %d, got %d", 3, i)
	}
}

func BenchmarkBinarySearch_BuiltinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchStrings(stringDS[:], "abc")
	}
}
