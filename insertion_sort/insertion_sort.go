package insertion_sort

import (
	"github.com/oleksiivelychko/go-code-helpers/array"
	"github.com/oleksiivelychko/go-code-helpers/intertype"
)

/*
InsertionSort O(n^2)

1. Find the smallest element in array.
2. Put the smallest element into new array.
3. Remove the smallest element from array.
*/
func InsertionSort[T intertype.INumber](slice []T) []T {
	var newSlice = make([]T, len(slice))

	var length = len(slice)
	for i := 0; i < length; i++ {
		smallestIndex := findSmallestIndex(slice)
		newSlice[i] = slice[smallestIndex]
		slice = array.Pop(slice, smallestIndex)
	}

	return newSlice
}

func findSmallestIndex[T intertype.INumber](slice []T) int {
	var smallestIndex = 0
	var smallest = slice[0]

	for i := 1; i <= len(slice)-1; i++ {
		if slice[i] < smallest {
			smallestIndex = i
			smallest = slice[i]
		}
	}

	return smallestIndex
}
