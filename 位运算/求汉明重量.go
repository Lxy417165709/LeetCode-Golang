package main

func hammingWeight(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1) // num = num & (num - 1) 会消除num最后一位的1
	}
	return count
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-1-bits/submissions/		位1的个数
*/

/*
	总结
	1. 这个题函数也适用于求负数1的个数。(补码)
*/
