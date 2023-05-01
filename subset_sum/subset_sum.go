package subset_sum

/*
SubsetSum - NP-complete task, does not exist efficient solution.
*/
func SubsetSum(set []int, n, sum int) bool {
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	include := SubsetSum(set, n-1, sum-set[n])
	exclude := SubsetSum(set, n-1, sum)

	return include || exclude
}
