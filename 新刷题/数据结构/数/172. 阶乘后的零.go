package 数

// trailingZeroes 获取 阶乘 中 因子10 的个数。
// 因子10 可以拆分为: 2、5。
func trailingZeroes(n int) int {
	return min(getCountOfFactorFromFactorial(n, 2), getCountOfFactorFromFactorial(n, 5))
}

// getCountOfFactorFromFactorial 获取 factorial阶乘 中 factor因子 个数。
func getCountOfFactorFromFactorial(factorial int, factor int) int {
	count := 0
	for factorial != 0 {
		count += factorial / factor
		factorial /= factor
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
