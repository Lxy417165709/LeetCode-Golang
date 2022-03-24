package 数

// climbStairs 获取爬到第n阶的方法数，一次可以爬 1步 or 2步。
func climbStairs(n int) int {
	// 1. 定义。
	dp := make([]int, 50) // dp[i] 表示爬到第i阶的方法数。

	// 2. 初始化。
	dp[1] = 1
	dp[2] = 2

	// 3. 动态规划。
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 4. 返回。
	return dp[n]
}
