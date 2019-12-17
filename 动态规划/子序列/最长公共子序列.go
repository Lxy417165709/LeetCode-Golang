package main

// 动态规划 解决最长公共子序列问题
// dp[i][t] 表示: text1[:i]与text2[:t]的最长公共子序列长度
// 状态转移方程:
//		i == 0 || t == 0		: dp[i][t] = 0
//		text1[i-1] == text2[t-1]: dp[i][t] = dp[i-1][t-1]+1
//      text1[i-1] != text2[t-1]: dp[i][t] = max(dp[i-1][t],dp[i][t-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := [1005][1005]int{}
	for i := 1; i <= len(text1); i++ {
		for t := 1; t <= len(text2); t++ {
			if text1[i-1] == text2[t-1] {
				dp[i][t] = dp[i-1][t-1] + 1
			} else {
				dp[i][t] = max(dp[i-1][t], dp[i][t-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-common-subsequence/		最长公共子序列
*/

/*
	总结
	1. 最长公共子序列和前面的子序列有点不一样，因为它涉及到了2个数组(字符串是一种字符数组)，
	   而前面的子序列只是涉及到1个数组。
*/
