package bubblesort

// sort O(n^2)
func sort(ds []int) []int {
	for i := 0; i < len(ds); i++ {
		for j := 0; j < len(ds)-i-1; j++ {
			if ds[j] > ds[j+1] {
				ds[j], ds[j+1] = ds[j+1], ds[j]
			}
		}
	}

	return ds
}
