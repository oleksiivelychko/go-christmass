package binarysearch

type integer interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type double interface {
	float32 | float64
}

type number interface {
	integer | double
}

// search O(log n)
// log2(8)=3 attempts are required to find value (operation opposite to exponentiation).
// Dataset must be sorted and contains numbers only.
func search[T number](ds []T, s T) (bool, int) {
	var bottom = 0
	var top = len(ds) - 1

	for bottom < top {
		// pick the middle item
		middle := (top + bottom) / 2

		// compare it to search value
		if ds[middle] == s {
			return true, middle
		} else if ds[middle] > s {
			// keep searching on the left side of dataset
			top = middle - 1
		} else if ds[middle] < s {
			// keep searching on the right side of dataset
			bottom = middle + 1
		}
	}

	return false, 0
}

func searchRecursion[T number](ds []T, s T) bool {
	var bottom = 0
	var top = len(ds) - 1
	var middle = -1
	var found = false

	if bottom < top {
		middle = (bottom + top) / 2

		if ds[middle] == s {
			return true
		} else if ds[middle] > s {
			top = middle - 1
		} else if ds[middle] < s {
			bottom = middle + 1
		}

		found = searchRecursion(ds[bottom:top+1], s)
	}

	return found
}
