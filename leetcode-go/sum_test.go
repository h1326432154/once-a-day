package leetcode

import "testing"

var (
	nums       = []int{2, 7, 7, 3, 6, 11, 15}
	target int = 9
	l          = len(nums)
)

/*
两数之和
问题描述:
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

// TestSumOfTwoNumber 解法1双重for
func TestSumOfTwoNumber(t *testing.T) {
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				t.Logf("值1为 %d,下标为 %d ,值2为 %d,下标为 %d", nums[i], i, nums[j], j)
			}
		}
	}
}
