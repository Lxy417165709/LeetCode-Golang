package 一维数组

func maxProfit(prices []int) int {
	// 1. 获取映射，leftMinValue[i] 表示 prices[:i+1] 的最小值。
	leftMinValue := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			leftMinValue[i] = prices[i]
			continue
		}
		leftMinValue[i] = min(leftMinValue[i-1], prices[i])
	}

	// 2. 获取映射，rightMaxValue[i] 表示 prices[i:] 的最大值。
	rightMaxValue := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			rightMaxValue[i] = prices[i]
			continue
		}
		rightMaxValue[i] = max(rightMaxValue[i+1], prices[i])
	}

	// 3. 获取最大收益。
	maxProfitValue := 0
	for i := 0; i < len(prices); i++ {
		maxProfitValue = max(maxProfitValue, rightMaxValue[i]-leftMinValue[i])
	}

	// 4. 返回。
	return maxProfitValue
}

func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
