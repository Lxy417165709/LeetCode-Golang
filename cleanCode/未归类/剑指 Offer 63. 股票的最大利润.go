package 未归类

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPriceBeforeNow := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		nowPrice := prices[i]
		minPriceBeforeNow = min(minPriceBeforeNow, prices[i])
		maxProfit = max(maxProfit, nowPrice-minPriceBeforeNow)
	}
	return maxProfit
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结
		1. 这题的抽象模型就是: 找出数组 2 个元素，A、B (A和B位于同一位置 或 A 在 B 之前)，求 B-A 的最大值。
*/
