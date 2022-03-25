package 字符串

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

// longestCommonSubsequence 获取 text1、text2 的最长公共子串长度。
func longestCommonSubsequence(text1 string, text2 string) int {
	// 1. 定义。
	dp := [1001][1001]int{} // 定义 dp[pos1][pos2] 为 text1[:pos1] 与 text2[:pos2] 的最长公共子串长度。

	// 2. 动态规划。
	for i := 0; i < len(text1); i++ {
		for t := 0; t < len(text2); t++ {
			pos1 := i + 1
			pos2 := t + 1
			if text1[i] == text2[t] {
				dp[pos1][pos2] = math_util.Max(dp[pos1-1][pos2-1]+1, dp[pos1-1][pos2], dp[pos1][pos2-1])
			} else {
				dp[pos1][pos2] = math_util.Max(dp[pos1-1][pos2], dp[pos1][pos2-1])
			}
		}
	}

	// 3. 返回。
	return dp[len(text1)][len(text2)]
}

