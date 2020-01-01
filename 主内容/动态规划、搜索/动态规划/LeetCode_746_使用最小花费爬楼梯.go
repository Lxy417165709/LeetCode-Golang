package 动态规划

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}



// 滚动优化
func minCostClimbingStairs(cost []int) int {
	// dp := make([]int,len(cost)+1)
	dp := make([]int, 2)

	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i&1] = min(dp[0], dp[1]) + cost[i]
	}
	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这题是简单的动态规划，和爬楼梯很像。
*/
