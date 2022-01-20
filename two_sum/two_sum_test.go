package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	sums := TwoSum([]int{2, 7, 11, 15}, 9)
	if sums[0] != 0 || sums[1] != 1 {
		t.Errorf("[func TwoSum(nums []int, target int) []int] -> [%d,%d] != [0,1]", sums[0], sums[1])
	}
	sums = TwoSum([]int{3, 2, 4}, 6)
	if sums[0] != 1 || sums[1] != 2 {
		t.Errorf("[func TwoSum(nums []int, target int) []int] -> [%d,%d] != [1,2]", sums[0], sums[1])
	}
	sums = TwoSum([]int{3, 3}, 6)
	if sums[0] != 0 || sums[1] != 1 {
		t.Errorf("[func TwoSum(nums []int, target int) []int] -> [%d,%d] != [0,1]", sums[0], sums[1])
	}
}
