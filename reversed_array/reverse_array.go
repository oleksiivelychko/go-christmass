package reversed_array

func reversedArray(dataset []int) []int {
	for i, j := 0, len(dataset)-1; i < j; i, j = i+1, j-1 {
		dataset[i], dataset[j] = dataset[j], dataset[i]
	}

	return dataset
}
