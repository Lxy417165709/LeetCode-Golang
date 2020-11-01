package main

/*
	两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
	给出两个整数 x 和 y，计算它们之间的汉明距离。

	注意：
	0 ≤ x, y < 2^31.
*/

// 汉明距离
func hammingDistance(x int, y int) int {
	return hammingWeight(x^y)
}
// 汉明重量
func hammingWeight(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

/*
	题目链接:
		https://leetcode-cn.com/problems/hamming-distance/		汉明距离
*/

/*
	总结
	1. 这题的思路是: 先获取x,y不同的比特位(通过异或运算)，之后求这个异或结果有多少个1就可以了。
*/