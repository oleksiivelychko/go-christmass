package in_array

import "testing"

var array = []string{"one", "two", "three"}

/*
*
-________|__Array_|__List__|
| Read   |  O(1)  |  O(n)  |
| Write  |  O(n)  |  O(1)  |
| Remove |  O(n)  |  O(1)  |
*/
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
