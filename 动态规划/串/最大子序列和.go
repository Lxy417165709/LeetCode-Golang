package main

const (
	inf = 100000000000 // 要定义的足够大，一般要大于int32
)

// 动态规划 解决最大子序和问题
// dp[i] 表示: 以nums[i-1]结尾能达到的最大和
// 状态转移方程:
// 			i == 0: dp[i] = 0
//			i >= 1: dp[i] = max(nums[i-1], dp[i-1]+nums[i-1])
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 1)
	ans := -inf
	for i := 1; i <= len(nums); i++ {
		dp = append(dp, max(nums[i-1], dp[i-1]+nums[i-1]))
		ans = max(ans, dp[i])
	}
	return ans
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], max(arr[1:]...)
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/maximum-subarray/submissions/
*/
