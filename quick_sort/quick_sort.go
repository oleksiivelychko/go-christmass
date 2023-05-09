package sort

import "github.com/oleksiivelychko/go-code-helpers/intertype"

/*
QuickSort

Best case (pivot is the middle) is O(n) * O(log n) = O(n log n)
Worst case (pivot is the first) is O(n) * O(n) = O(n^2)

1. Select item as pivot.
2. Compare items - all items less than pivot value move before it, all items greater than pivot move after it.
3. Repeat steps 1,2 towards arrays - left and right between pivot.
*/
func QuickSort[T intertype.INumber](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	pivot := slice[0]

	var head []T
	for _, num := range slice[1:] {
		if num <= pivot {
			head = append(head, num)
		}
	}

	var tail []T
	for _, num := range slice[1:] {
		if num > pivot {
			tail = append(tail, num)
		}
	}

	head = QuickSort(head)
	tail = QuickSort(tail)

	head = append(head, pivot)
	head = append(head, tail...)

	return head
}
