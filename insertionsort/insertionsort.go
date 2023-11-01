package insertionsort

// sort O(n^2)
func sort(ds []int) []int {
	var sorted = make([]int, len(ds))

	for i := 0; i < len(ds); i++ {
		// get smallest item in dataset
		smallest := findSmallest(ds)
		// put smallest item into dataset
		sorted[i] = ds[smallest]
		// pop smallest item from dataset by index
		ds = append(ds[:smallest], ds[smallest+1:]...)
	}

	return sorted
}

func findSmallest(dataset []int) int {
	var index = 0
	var smallest = dataset[0]

	for i := 1; i <= len(dataset)-1; i++ {
		if dataset[i] < smallest {
			index = i
			smallest = dataset[i]
		}
	}

	return index
}
