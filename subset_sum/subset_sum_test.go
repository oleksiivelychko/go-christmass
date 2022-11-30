package subset_sum

import "testing"

func TestSubsetSum(t *testing.T) {
	set := []int{3, 4, 5, 2}
	isSubsetSum := SubsetSum(set, len(set)-1, 9)
	if isSubsetSum != true {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != true", isSubsetSum)
	}
}
