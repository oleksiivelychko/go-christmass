package binary_search

type IInt interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type IFloat interface {
	float32 | float64
}

type INumber interface {
	IInt | IFloat
}

/*
BinarySearch O(log n)
log2(8)=3 attempts are required to find value (operation opposite to exponentiation).

Dataset must be sorted and contains numbers only.

1. pick the middle item and compare it to search value
2. if picked value less than search one then keep searching on the right side of dataset
3. if picked value greater than search one then keep searching on the left side of dataset
4. repeat steps until search value will not be found or dataset is not empty
*/
func BinarySearch[T INumber](slice []T, target T) (bool, int) {
	var bottom = 0
	var top = len(slice) - 1

	for bottom < top {
		middle := (top + bottom) / 2

		if slice[middle] == target {
			return true, middle
		} else if slice[middle] > target {
			top = middle - 1
		} else if slice[middle] < target {
			bottom = middle + 1
		}
	}

	return false, 0
}

func BinarySearchRecursion[T INumber](slice []T, target T) bool {
	var bottom = 0
	var top = len(slice) - 1
	var middle = -1
	var found = false

	if bottom < top {
		middle = (bottom + top) / 2

		if slice[middle] == target {
			return true
		} else if slice[middle] > target {
			top = middle - 1
		} else if slice[middle] < target {
			bottom = middle + 1
		}

		found = BinarySearchRecursion(slice[bottom:top+1], target)
	}

	return found
}
