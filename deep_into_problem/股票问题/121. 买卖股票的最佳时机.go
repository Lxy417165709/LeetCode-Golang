package 股票问题

// ---------------------- 朴素 DP ----------------------
// 第 i 天卖出的最大利润 = 第 i 天股票的价格 - 前 i-1 天股票的最低价格(此时购入股票)
// 即: maxProfitInDay[i] = prices[i]-min(prices[:i]...)
// 下面的代码会超时，但可以用前缀和优化 min(prices[:i]...)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfitInDay := make([]int, len(prices))
	maxProfitInDay[0] = 0
	for i := 1; i < len(prices); i++ {
		maxProfitInDay[i] = prices[i] - min(prices[:i]...)
	}
	return max(maxProfitInDay...)
}

// ---------------------- 朴素 DP + 前缀和优化----------------------
// 优化: 去除前缀和数组
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfitInDay := make([]int, len(prices))

	minPrice := make([]int, len(prices))
	minPrice[0] = prices[0]
	for i := 1; i < len(minPrice); i++ {
		minPrice[i] = min(minPrice[i-1], prices[i])
	}

	maxProfitInDay[0] = 0
	for i := 1; i < len(prices); i++ {
		maxProfitInDay[i] = prices[i] - minPrice[i-1]
	}
	return max(maxProfitInDay...)
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], max(arr[1:]...)
	if a > b {
		return a
	}
	return b
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}
