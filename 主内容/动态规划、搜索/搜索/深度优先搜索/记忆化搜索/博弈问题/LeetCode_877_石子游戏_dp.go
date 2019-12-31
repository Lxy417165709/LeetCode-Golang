package main

// dp解决
func stoneGame(piles []int) bool {
	if len(piles) == 0 {
		return true
	}
	sum := [505]int{} // sum[t+1]-sum[i] 表示 piles[i:t]的石头堆总数
	for i := 1; i <= len(piles); i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}

	// dp[i][t][0] 表示先手遇到[i,t]局势所能获得的最多石头
	// dp[i][t][1] 表示后手遇到[i,t]局势所能获得的最多石头
	dp := [505][505][2]int{}

	for i := len(piles) - 1; i >= 0; i-- {
		for t := 0; t < len(piles); t++ {
			if i > t {
				continue
			}
			if i == t {
				dp[i][t][0] = piles[i]
				dp[i][t][1] = 0
				continue
			}
			left := piles[i] + dp[i+1][t][1]
			right := piles[t] + dp[i][t-1][1]
			dp[i][t][0] = max(left, right)
			dp[i][t][1] = sum[t+1] - sum[i] - dp[i][t][0]
		}
	}
	// fmt.Println(dp[0][len(piles)-1][0],dp[0][len(piles)-1][1])
	return dp[0][len(piles)-1][0] > dp[0][len(piles)-1][1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 状态压缩
func stoneGame(piles []int) bool {
	if len(piles) == 0 {
		return true
	}
	sum := [505]int{} // sum[t+1]-sum[i] 表示 piles[i:t]的石头堆总数
	for i := 1; i <= len(piles); i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}


	dp := [505][2]int{}

	for i := len(piles) - 1; i >= 0; i-- {
		for t := 0; t < len(piles); t++ {
			if i > t {
				continue
			}
			if i == t {
				dp[t][0] = piles[i]
				dp[t][1] = 0
				continue
			}
			left := piles[i] + dp[t][1]
			right := piles[t] + dp[t-1][1]
			dp[t][0] = max(left, right)
			dp[t][1] = sum[t+1] - sum[i] - dp[t][0]
		}
	}
	// fmt.Println(dp[0][len(piles)-1][0],dp[0][len(piles)-1][1])
	return dp[len(piles)-1][0] > dp[len(piles)-1][1]
}

// 最终版本
func stoneGame(piles []int) bool {
	if len(piles) == 0 {
		return true
	}
	sum := [505]int{} // sum[t+1]-sum[i] 表示 piles[i:t]的石头堆总数
	for i := 1; i <= len(piles); i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}


	dp := [505][2]int{}

	for i := len(piles) - 1; i >= 0; i-- {
		dp[i][0] = piles[i]
		dp[i][1] = 0
		for t := i+1; t < len(piles); t++ {
			left := piles[i] + dp[t][1]
			right := piles[t] + dp[t-1][1]
			dp[t][0] = max(left, right)
			dp[t][1] = sum[t+1] - sum[i] - dp[t][0]
		}
	}
	return dp[len(piles)-1][0] > dp[len(piles)-1][1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
	总结
	1. 注意: 以上代码提交会超时！
	2. 实际上这道题要采用dp解，这样就不会超时了。
*/
