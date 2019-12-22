package main

func subtractProductAndSum(n int) int {
	if n < 0 {
		n = -n
	}
	mul, sum := 1, 0
	for n != 0 {
		mul *= n % 10
		sum += n % 10
		n /= 10
	}
	return mul - sum
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/		整数的各位积和之差
*/
/*
	总结
	1. 这是一道超级水的水题。
*/
