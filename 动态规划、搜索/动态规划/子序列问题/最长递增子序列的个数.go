package main

// dp[i] 表示: 		以nums[i]结尾的最长递增子序列长度
// count[i]表示: 	以nums[i]结尾的最长递增子序列个数
func findNumberOfLIS(nums []int) int {
	dp := [2005]int{}
	count := make(map[int]int, len(nums)+5)
	maxLength := 0 // nums中最长递增子序列长度
	for i := 0; i < len(nums); i++ {
		dp[i] = 1		// 注意初始化
		count[i] = 1	// 注意初始化
		for t := 0; t < i; t++ {
			if nums[i] > nums[t] {
				if dp[t]+1 == dp[i] {
					count[i] += count[t]
				} else {
					if dp[t]+1 > dp[i] {
						dp[i] = dp[t] + 1
						count[i] = count[t]
					}
				}
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLength {
			ans += count[i]

		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/		最长递增子序列的个数
*/
/*
	总结
	1. 之前就在纠结，当nums[i]==nums[t]时是否要拓展，纠结好久后就去看答案了，发现等于时不用拓展。
	   只在nums[i] > nums[t]时拓展。

*/