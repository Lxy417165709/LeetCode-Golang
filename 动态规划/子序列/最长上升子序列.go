package main

// 动态规划 解决最大子序和问题
// dp[i] 表示: 以nums[i]结尾的最长上升子序列长度
// 状态转移方程:
//		i == 0:				dp[i] = 1
// 		nums[i] > nums[t]:	dp[i] = max(1, dp[i], dp[t]+1)		( t∈[0, i) )
func lengthOfLIS(nums []int) int {
	dp := make([]int, 5005)
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for t := 0; t < i; t++ {
			if nums[i] > nums[t] {
				dp[i] = max(dp[i], dp[t]+1)
			}
		}
		ans = max(dp[i], ans)
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
	题目链接: https://leetcode-cn.com/problems/longest-increasing-subsequence/submissions/
*/
