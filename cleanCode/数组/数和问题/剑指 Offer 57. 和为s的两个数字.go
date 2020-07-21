package 数和问题

import "sort"

func twoSum(nums []int, target int) []int {
	return pairSums(nums, target)[0]
}

func pairSums(nums []int, target int) [][]int {
	pairs := make([][]int, 0)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			pairs = append(pairs, []int{nums[left], nums[right]})
			left++
			right--
		} else {
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return pairs
}

/*
	题目链接: https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
*/



