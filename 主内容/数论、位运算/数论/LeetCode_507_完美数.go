package main

func checkPerfectNumber(num int) bool {
	return num > 0 && getFactorSum(num)-num == num
}

// 获取num的不同的因子和 ( 时间复杂度O(n^(1/2)) )
func getFactorSum(x int) int {
	sum := 0
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			sum += i
			if x/i != i {
				sum += x / i
			}
		}
	}
	return sum
}

/*
	题目链接:
		https://leetcode-cn.com/problems/perfect-number/submissions/		完美数
*/

/*
	总结
	1. 上面的解答采用了常规解法，官方还提供了欧几里得解法..把时间复杂度优化为O(1)
*/
