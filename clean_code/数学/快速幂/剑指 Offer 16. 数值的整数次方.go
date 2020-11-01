package 快速幂

func myPow(x float64, n int) float64 {
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	isExponentNegative := n < 0
	notNegativeExponent := abs(n)
	result, curWeight := 1.0, x
	for notNegativeExponent != 0 {
		if (notNegativeExponent & 1) == 1 {
			result *= curWeight
		}
		curWeight *= curWeight
		notNegativeExponent >>= 1
	}
	if isExponentNegative {
		return 1.0 / result
	}
	return result
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接: https://leetcode-cn.com/problems/powx-n/
	总结:
		1. 这题要注意的就是: 指数可能是负数
*/
