package 股票问题

// 状态: 所处天、是否有股票、交易数
// 操作: 什么都不做、买入、卖出
// 定义: 处于某状态能得到的最大利益
// dp[i][0][0] = max(dp[i-1][0][0],dp[i-1][1][-1]+prices[i])
// dp[i][0][1] = max(dp[i-1][0][1],dp[i-1][1][0]+prices[i])
// dp[i][0][2] = max(dp[i-1][0][2],dp[i-1][1][1]+prices[i])
// dp[i][1][0] = max(dp[i-1][1][0],dp[i-1][0][0]-prices[i])
// dp[i][1][1] = max(dp[i-1][1][1],dp[i-1][0][1]-prices[i])
// dp[i][1][2] = max(dp[i-1][1][2],dp[i-1][0][2]-prices[i])

const inf = 1000000000

func maxProfit(prices []int) int {
	maxProfitResult := make(map[int]int)
	hasStockStatus := []int{0, 1}
	countOfSellStatus := []int{0, 1, 2}
	result := 0
	for day, price := range prices {
		for _, hasStock := range hasStockStatus {
			for _, countOfSell := range countOfSellStatus {
				if hasStock == 0 {
					if _, ok := maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)]; !ok {
						maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)] = -inf
					}
					maxProfitResult[hash(day, hasStock, countOfSell)] = max(
						maxProfitResult[hash(day-1, hasStock, countOfSell)],
						maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)]+price,
					)
				}else{
					maxProfitResult[hash(day, hasStock, countOfSell)] = max(
						maxProfitResult[hash(day-1, hasStock, countOfSell)],
						maxProfitResult[hash(day-1, hasStock^1, countOfSell)]-price,
					)
				}

				result = max(result, maxProfitResult[hash(day, hasStock, countOfSell)])
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hash(day, hasStock, countOfSell int) int {
	return day*100000 + hasStock*10 + countOfSell
}
