package main

// 递归 + 位运算 实现加法
func getSum(a int, b int) int {
	num1 := (a & b) << 1
	num2 := a ^ b

	if num1 == 0 {
		return num2
	}
	return getSum(num1, num2)
}

// 迭代 + 位运算 实现加法
func getSum(a int, b int) int {
	for {
		num1 := (a & b) << 1
		num2 := a ^ b
		if num1 == 0 {
			return num2
		}
		a, b = num1, num2
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/sum-of-two-integers/
*/

/*
	总结
	1. 思路:
			(1) a + b == a与b的不进位加法值 + a与b相加的进位值
			(2) 其中 a与b的不进位加法值 == a^b，a与b相加的进位值 == (a & b) << 1
			(3) 循环这个过程:
					直到进位值 a与b相加的进位值 == 0，此时返回 a与b的不进位加法值。

*/
