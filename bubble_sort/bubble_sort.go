package sort

/*
BubbleSort O(n^2)
*/
func BubbleSort(dataset []int) []int {
	for i := 0; i < len(dataset); i++ {
		for j := 0; j < len(dataset)-i-1; j++ {
			if dataset[j] > dataset[j+1] {
				dataset[j], dataset[j+1] = dataset[j+1], dataset[j]
			}
		}
	}
	return dataset
}
