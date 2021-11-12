package binary_search

/*
BinarySearch

O(log n)
log2(8)=3 attempts required to find value (operation opposite to exponentiation).

Array must be sortable before.

1. Pick middle item.
2. If middle item less than value then search will be in left side of array.
3. If middle item greater than value then search will be in right side of array.
4. Repeat steps until target won't be found.
*/
func BinarySearch(array []int, target int) (bool, int) {

	var bottom = 0
	var top = len(array) - 1

	for bottom < top {

		var middle = bottom + top

		if array[middle] == target {
			return true, middle
		}

		if array[middle] > target {
			top = middle - 1
		}

		if array[middle] < target {
			bottom = middle + 1
		}
	}

	return false, 0
}

func BinarySearchRecursion(array []int, target int) (bool, int) {

	var bottom = 0
	var top = len(array) - 1
	var middle = -1

	if bottom < top {
		middle = bottom + top

		if array[middle] == target {
			return true, middle
		}

		if array[middle] > target {
			top = middle - 1
		}

		if array[middle] < target {
			bottom = middle + 1
		}

		_, middle = BinarySearchRecursion(array[bottom:top+1], target)
	}

	return false, middle
}
