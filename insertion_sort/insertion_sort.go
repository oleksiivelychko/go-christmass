package insertion_sort

/*
InsertionSort O(n^2)
*/
func InsertionSort(dataset []int) []int {
	var sortedDataset = make([]int, len(dataset))

	for i := 0; i < len(dataset); i++ {
		// get smallest item in dataset
		smallestIndex := findSmallest(dataset)
		// put smallest item into dataset
		sortedDataset[i] = dataset[smallestIndex]
		// pop smallest item from dataset by index
		dataset = append(dataset[:smallestIndex], dataset[smallestIndex+1:]...)
	}

	return sortedDataset
}

func findSmallest(dataset []int) int {
	var smallestIndex = 0
	var smallest = dataset[0]

	for i := 1; i <= len(dataset)-1; i++ {
		if dataset[i] < smallest {
			smallestIndex = i
			smallest = dataset[i]
		}
	}

	return smallestIndex
}
