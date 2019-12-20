package main

import "sort"

/*
	给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1),
	(a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。
*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)>>1; i++ {
		ans += nums[i<<1]
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/array-partition-i/			数组拆分 I
*/
