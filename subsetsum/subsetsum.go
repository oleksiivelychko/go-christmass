package subsetsum

// subsetSum - NP-complete task, does not exist efficient solution.
func subsetSum(ds []int, n, sum int) bool {
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	include := subsetSum(ds, n-1, sum-ds[n])
	exclude := subsetSum(ds, n-1, sum)

	return include || exclude
}
