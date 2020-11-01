package Geek

func checkPossibility(nums []int) bool {
	timesOfChanging := 0
	nums = append([]int{-1000000000000}, nums...)
	for i := 2; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			timesOfChanging++
			if nums[i] >= nums[i-2] {
				nums[i-1] = nums[i-2]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return timesOfChanging <= 1
}

/*
	题目链接: https://leetcode-cn.com/problems/non-decreasing-array/
	总结
		1. 这题思路还没理清，2020年7月5日03:40:47。
*/
