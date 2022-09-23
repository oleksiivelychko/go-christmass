package reverse_array

import (
	"testing"
)

/**
-________|__Array_|__List__|
| Read   |  O(1)  |  O(n)  |
| Write  |  O(n)  |  O(1)  |
| Remove |  O(n)  |  O(1)  |
*/

func makeSlice() []int {
	var slice = make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	return slice
}

func TestSmallestSearch(t *testing.T) {
	reversedArray := ReverseArray(makeSlice())
	if reversedArray[0] != 5 {
		t.Errorf("[func ReverseArray(slice []int) []int] -> %d != -2", reversedArray[0])
	}
	if reversedArray[1] != 4 {
		t.Errorf("[func ReverseArray(slice []int) []int] -> %d != -2", reversedArray[1])
	}
	if reversedArray[2] != 3 {
		t.Errorf("[func ReverseArray(slice []int) []int] -> %d != -2", reversedArray[2])
	}
	if reversedArray[3] != 2 {
		t.Errorf("[func ReverseArray(slice []int) []int] -> %d != -2", reversedArray[3])
	}
	if reversedArray[4] != 1 {
		t.Errorf("[func ReverseArray(slice []int) []int] -> %d != -2", reversedArray[4])
	}
}

func BenchmarkSmallestSearch(b *testing.B) {
	slice := makeSlice()
	for i := 0; i < b.N; i++ {
		ReverseArray(slice)
	}
}
