package insertion_sort

/*
InsertionSort O(n^2)

1. Find the smallest element in array.
2. Put the smallest element into new array.
3. Remove the smallest element from array.
*/
func InsertionSort(slice []int) []int {
	var newSlice = make([]int, len(slice))

	var length = len(slice)
	for i := 0; i < length; i++ {
		smallestIndex := findSmallestIndex(slice)
		newSlice[i] = slice[smallestIndex]
		// pop element from array by index
		slice = append(slice[:smallestIndex], slice[smallestIndex+1:]...)
	}

	return newSlice
}

func findSmallestIndex(slice []int) int {
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
