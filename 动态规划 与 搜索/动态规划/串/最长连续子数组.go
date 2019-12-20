package main

// dp[i][t] 表示: 以A[i-1]、B[t-1]结尾的公共数组长度(从右到左可以拓展多长)
// 状态转移方程:
// 			A[i-1] == B[t-1]: dp[i][t] = dp[i-1][t-1] + 1
//			A[i-1] != B[t-1]: dp[i][t] = 0
func findLength(A []int, B []int) int {
	dp := [1005][1005]int{}
	ans := 0
	for i := 1; i <= len(A); i++ {
		for t := 1; t <= len(B); t++ {
			if A[i-1] == B[t-1] {
				dp[i][t] = dp[i-1][t-1] + 1
			} else {
				dp[i][t] = 0
			}
			ans = max(ans, dp[i][t])
		}
	}
	return ans
}

// 滚动优化
func findLength(A []int, B []int) int {
	dp := [1005]int{}
	ans := 0
	for i := 1; i <= len(A); i++ {
		// 注意此时需要逆序遍历，保证结果不会被覆盖
		for t := len(B); t >= 1; t-- {
			if A[i-1] == B[t-1] {
				dp[t] = dp[t-1] + 1
			} else {
				dp[t] = 0
			}
			ans = max(ans, dp[t])
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
		https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/		最长重复子数组
*/
/*
	总结
	1. 这题和最长相同子序列是不一样的，序列可以不连续，而这题是需要连续的。
	2. 官方还使用了二分法做...
*/
