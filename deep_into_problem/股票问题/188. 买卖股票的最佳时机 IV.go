package 股票问题

const inf = 1000000000

func maxProfit(k int, prices []int) int {
	maxProfitResult := make(map[int]int)
	hasStockStatus := []int{0, 1}
	result := 0
	for day, price := range prices {
		for _, hasStock := range hasStockStatus {
			for countOfSell := 0; countOfSell <= k; countOfSell++ {
				if hasStock == 0 {
					if _, ok := maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)]; !ok {
						maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)] = -inf
					}
					maxProfitResult[hash(day, hasStock, countOfSell)] = max(
						maxProfitResult[hash(day-1, hasStock, countOfSell)],
						maxProfitResult[hash(day-1, hasStock^1, countOfSell-1)]+price,
					)
				} else {
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
	return day*10000000000000 + hasStock*1000000001 + countOfSell
}

/*
可以参考 123 的状态转移
*/
