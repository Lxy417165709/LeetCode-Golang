package 动态规划

const INF = 100000000000000000

func coinChange(coins []int, amount int) int {
	minCoinsOfFormingSum := make([]int, amount+1)
	makeSliceINF(minCoinsOfFormingSum)
	minCoinsOfFormingSum[0] = 0
	for i := 1; i <= amount; i++ {
		for t := 0; t < len(coins); t++ {
			if i-coins[t] >= 0 {
				minCoinsOfFormingSum[i] = min(minCoinsOfFormingSum[i], minCoinsOfFormingSum[i-coins[t]]+1)
			}
		}
	}
	if minCoinsOfFormingSum[amount] == INF {
		return -1
	}
	return minCoinsOfFormingSum[amount]
}

func makeSliceINF(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = INF
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


/*
	题目链接: https://leetcode-cn.com/problems/coin-change/submissions/
	总结
		1. 如果数据量比较小，其实可以使用排列组合AC这题。
*/
