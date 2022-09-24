package in_array

import "testing"

var array = []string{"one", "two", "three"}

func TestInArray(t *testing.T) {
	index, result := InArray(array, "one")
	if !result {
		t.Errorf("[func InArray[C comparable](array []C, needle C) (int, bool)] -> %d != '0'", index)
	}
}

func BenchmarkInArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InArray(array, "one")
	}
}
