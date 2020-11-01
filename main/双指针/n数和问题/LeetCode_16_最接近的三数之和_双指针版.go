package main

import "sort"

func threeSumClosest(nums []int, target int) int {

	// 题目限制了只有唯一答案，所以也说明了len(nums)>=3
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	minDistance := 100000000000
	ans := -1
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if abs(sum-target) < minDistance {
				minDistance = abs(sum - target)
				ans = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}

		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	总结
	1. 这题之前我是用暴力+二分搜索AC的，代码很长也有些乱。
	2. 采用双指针做的时候代码更简洁了，时间复杂度也更低了
*/
