package two_sum

/*
TwoSum
https://leetcode.com/problems/two-sum/
*/
func TwoSum(nums []int, target int) []int {
	sums := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sums = append(sums, i, j)
				return sums
			}
		}
	}
	return sums
}
