package 股票问题

// ------------------ 单调栈法 ------------------
// 使用单调递减栈，找 A[:i]从右到左第一个小于A[i]的数
// 优化: 单调栈可以优化掉...
func maxProfit(prices []int) int {
	minStack := make([]int, 0)
	maxProfitResult := 0
	for i := 0; i < len(prices); i++ {
		if len(minStack) != 0 && minStack[len(minStack)-1] <= prices[i] {
			maxProfitResult += prices[i] - minStack[len(minStack)-1]
			for len(minStack) != 0 && minStack[len(minStack)-1] <= prices[i] {
				minStack = minStack[:len(minStack)-1]
			}
		}
		minStack = append(minStack, prices[i])
	}
	return maxProfitResult
}

// 使用单调递增栈，找 A[i+1:]从左到右第一个小于A[i]的数
func maxProfit(prices []int) int {
	maxStack := make([]int, 0)
	maxProfitResult := 0
	for i := 0; i < len(prices); i++ {
		if len(maxStack) != 0 && maxStack[len(maxStack)-1] >= prices[i] {
			maxProfitResult += maxStack[len(maxStack)-1] - maxStack[0]
			maxStack = []int{}
		}
		maxStack = append(maxStack, prices[i])
	}
	if len(maxStack) != 0 {
		maxProfitResult += maxStack[len(maxStack)-1] - maxStack[0]
	}
	return maxProfitResult
}

// ------------------ 动态规划 ------------------
// 定义: 前 i 天能获得的最大利润
// 状态: 手上是否有股票
// 操作: 什么都不做、买入、卖出
// dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])


