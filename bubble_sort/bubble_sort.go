package sort

import "github.com/oleksiivelychko/go-code-helpers/intertype"

/*
BubbleSort O(n^2)
*/
func BubbleSort[T intertype.INumber](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
