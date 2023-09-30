package sort

/*
quickSort
Worst case (pivot is the first) is O(n) * O(n) = O(n^2)
Best case (pivot is the middle) is O(n) * O(log n) = O(n log n)
*/
func quickSort(dataset []int) []int {
	if len(dataset) == 0 {
		return dataset
	}

	pivot := dataset[0]

	var head []int
	for _, item := range dataset[1:] {
		// items that are smaller than pivot need to be moved before it
		if item <= pivot {
			head = append(head, item)
		}
	}

	var tail []int
	for _, item := range dataset[1:] {
		// items that are greater than pivot need to be moved after it
		if item > pivot {
			tail = append(tail, item)
		}
	}

	head = quickSort(head)
	tail = quickSort(tail)

	// pivot equals to the smallest value in dataset after the first recursive invoking
	head = append(head, pivot)
	head = append(head, tail...)

	return head
}
