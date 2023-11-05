package insertionsort

// sort O(n^2)
func sort(ds []int) []int {
	var (
		lenDS  = len(ds)
		sorted = make([]int, lenDS)
	)

	for i := 0; i < lenDS; i++ {
		// get smallest item from dataset
		smallest := findSmallest(ds)
		// put smallest item into dataset
		sorted[i] = ds[smallest]
		// pop smallest item from dataset by index
		ds = append(ds[:smallest], ds[smallest+1:]...)
	}

	return sorted
}

func findSmallest(ds []int) int {
	var (
		index    = 0
		smallest = ds[0]
	)

	for i := 1; i <= len(ds)-1; i++ {
		if ds[i] < smallest {
			index = i
			smallest = ds[i]
		}
	}

	return index
}
