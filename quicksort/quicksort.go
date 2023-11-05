package quicksort

// Worst case (pivot is the first) is O(n) * O(n) = O(n^2)
// Best case (pivot is the middle) is O(n) * O(log n) = O(n log n)
func sort(ds []int) []int {
	if len(ds) == 0 {
		return ds
	}

	var (
		pivot = ds[0]
		head  []int
		tail  []int
	)

	for _, item := range ds[1:] {
		// items that are smaller than pivot need to be moved before it
		if item <= pivot {
			head = append(head, item)
		}
	}

	for _, item := range ds[1:] {
		// items that are greater than pivot need to be moved after it
		if item > pivot {
			tail = append(tail, item)
		}
	}

	head = sort(head)
	tail = sort(tail)

	// pivot equals to the smallest value in dataset after the first recursive invoking
	head = append(head, pivot)
	head = append(head, tail...)

	return head
}
