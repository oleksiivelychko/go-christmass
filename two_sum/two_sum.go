package two_sum

/*
TwoSum
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Constraints:
2 <= nums.length <= 104
-10^9 <= nums[i] <= 109
-10^9 <= target <= 109
Only one valid answer exists.
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
