package 动态规划

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		minCost[i] = min(minCost[i-1]+cost[i-1], minCost[i-2]+cost[i-2])
	}
	return minCost[len(cost)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/min-cost-climbing-stairs/
*/
