package 字符串

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
				dp[pos1][pos2] = max(dp[pos1-1][pos2-1]+1, dp[pos1-1][pos2], dp[pos1][pos2-1])
			} else {
				dp[pos1][pos2] = max(dp[pos1-1][pos2], dp[pos1][pos2-1])
			}
		}
	}

	// 3. 返回。
	return dp[len(text1)][len(text2)]
}

// max 获取数组最大值。
func max(nums ...int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	return max(nums[0], max(nums[1:]...))
}
