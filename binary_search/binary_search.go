package binary_search

type iInt interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type iFloat interface {
	float32 | float64
}

type iNumber interface {
	iInt | iFloat
}

/*
binarySearch O(log n)
log2(8)=3 attempts are required to find value (operation opposite to exponentiation).
Dataset must be sorted and contains numbers only.
*/
func binarySearch[T iNumber](dataset []T, searchValue T) (bool, int) {
	var bottom = 0
	var top = len(dataset) - 1

	for bottom < top {
		// pick the middle item
		middle := (top + bottom) / 2

		// compare it to search value
		if dataset[middle] == searchValue {
			return true, middle
		} else if dataset[middle] > searchValue {
			// keep searching on the left side of dataset
			top = middle - 1
		} else if dataset[middle] < searchValue {
			// keep searching on the right side of dataset
			bottom = middle + 1
		}
	}

	return false, 0
}

func binarySearchRecursion[T iNumber](dataset []T, searchValue T) bool {
	var bottom = 0
	var top = len(dataset) - 1
	var middle = -1
	var found = false

	if bottom < top {
		middle = (bottom + top) / 2

		if dataset[middle] == searchValue {
			return true
		} else if dataset[middle] > searchValue {
			top = middle - 1
		} else if dataset[middle] < searchValue {
			bottom = middle + 1
		}

		found = binarySearchRecursion(dataset[bottom:top+1], searchValue)
	}

	return found
}
