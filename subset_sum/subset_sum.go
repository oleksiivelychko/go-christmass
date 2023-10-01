package subset_sum

/*
subsetSum - NP-complete task, does not exist efficient solution.
*/
func subsetSum(set []int, n, sum int) bool {
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	include := subsetSum(set, n-1, sum-set[n])
	exclude := subsetSum(set, n-1, sum)

	return include || exclude
}
