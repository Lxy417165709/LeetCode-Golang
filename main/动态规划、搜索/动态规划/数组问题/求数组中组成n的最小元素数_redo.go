package main

// 获取: 使用数组 nums 中的元素，组成 target 的最少元素数
// 前提: 数组元素大于0、不存在重复元素(存在也可以)、可重复选取
// dp[i]表示: 使用数组 nums 中的元素，组成 i 的最少元素数
func minimumQuantity(nums []int, target int) int {
	const inf  = 1000000000
	if target < 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 1; i <= target; i++ {
		dp[i] = inf
		for t := 0; t < len(nums); t++ {
			if i >= nums[t] {
				dp[i] = min(dp[i], dp[i-nums[t]]+1)
			}
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
/*
	322. 零钱兑换2: https://leetcode-cn.com/problems/coin-change/
*/
