package main

func minimumDeleteSum(word1 string, word2 string) int {
	dp := [1005][1005]int{} // dp[i][t] 表示 word1[:i] 转换为 word2[:t] 所需删除字符的ASCII值的最小和。

	/* 要初始化 */
	dp[0][0] = 0
	// 空处理
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = dp[i-1][0] + int(word1[i-1])
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = dp[0][i-1] + int(word2[i-1])
	}

	// 二者非空处理
	for i := 1; i <= len(word1); i++ {
		for t := 1; t <= len(word2); t++ {
			if word1[i-1] == word2[t-1] {
				dp[i][t] = dp[i-1][t-1]
			} else {
				// 和两个字符串的删除操作差不多
				dp[i][t] = min(dp[i-1][t]+int(word1[i-1]), dp[i][t-1]+int(word2[t-1]))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], min(nums[1:]...)
	if a > b {
		return b
	}
	return a
}

// 滚动优化 (有点复杂。。)
func minimumDeleteSum(word1 string, word2 string) int {
	dp := [1005]int{} // dp[t] 表示 word1 转换为 word2[:t] 所需删除字符的ASCII值的最小和。

	/* 要初始化 */
	dp[0] = 0

	// 当word1为空时
	for t := 1; t <= len(word2); t++ {
		dp[t] = dp[t-1] + int(word2[t-1])
	}
	// 当word1不为空时
	for i := 1; i <= len(word1); i++ {
		dp_i_1_t_1 := dp[0]	// 表示的是 dp[i-1][t-1]
		dp[0] += int(word1[i-1])
		for t := 1; t <= len(word2); t++ {
			dp_i_1 := dp[t]	// 表示 dp[i-1][t]
			if word1[i-1] == word2[t-1] {
				dp[t] = dp_i_1_t_1
			} else {
				// 与编辑距离相比，就这里少了一个操作
				dp[t] = min(dp[t]+int(word1[i-1]), dp[t-1]+int(word2[t-1]))
			}
			dp_i_1_t_1 = dp_i_1
		}
	}
	// 状态压缩: 二维 -> 一维 (这里可能不太清楚)
	// dp[i-1][t] 	: dp[t]
	// dp[i][t-1]   : dp[t-1]
	// dp[i-1][t-1]	: dp_i_1_t_1
	// dp[i-1][t]	: dp_i_1 (用于临时保存dp[t]，因为接下来dp[t]会改变)
	// 用二维dp模拟一下就可以知道这个状态了。

	return dp[len(word2)]
}

/*
	总结
	1. 这题和编辑距离是一样的，代码改一句话就可以了。
	2. 还可以用最长公共子序列解这个题。
*/
