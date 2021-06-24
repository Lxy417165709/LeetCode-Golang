package 股票问题



// 状态: 天数、股票拥有状态
// 操作: 什么都不做、买入、卖出
// 定义: 处于某状态下能获得的最大利益
// dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i]-fee)
// dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
const inf = 10000000000
func maxProfit(prices []int,fee int) int {
	maxProfitResult := make(map[int]int)
	result := 0

	maxProfitResult[hash(-1,1)] = -inf
	for day, price := range prices {
		maxProfitResult[hash(day,0)] = max(
			maxProfitResult[hash(day-1,0)],
			maxProfitResult[hash(day-1,1)]+price-fee,
		)
		maxProfitResult[hash(day,1)] = max(
			maxProfitResult[hash(day-1,1)],
			maxProfitResult[hash(day-1,0)]-price,
		)
		result = max(result, maxProfitResult[hash(day, 0)])
		result = max(result, maxProfitResult[hash(day, 1)])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hash(day, hasStock int) int {
	return day*100000 + hasStock*10
}

