package binary_search

import (
	"math/rand"
	"testing"
)

var array = [8]int{1, 3, 5, 7, 9, 11, 13, 15}

func TestBinarySearch(t *testing.T) {
	result, index := BinarySearch(array[:], 5)
	if !result && index != 2 {
		t.Errorf("[func BinarySearch(array []int, target int) (bool, int)] -> (%t, %d) != (true, 2)", result, index)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(array[:], rand.Intn(7))
	}
}

func TestBinarySearchRecursion(t *testing.T) {
	result, index := BinarySearchRecursion(array[:], 9)
	if !result && index != 4 {
		t.Errorf("[func BinarySearchRecursion(array []int, target int) (bool, int)] -> (%t, %d) != (true, 4)", result, index)
	}
}

func BenchmarkBinarySearchRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(array[:], rand.Intn(7))
	}
}
