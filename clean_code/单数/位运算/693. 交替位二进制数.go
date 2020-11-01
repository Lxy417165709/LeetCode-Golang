package 位运算

func hasAlternatingBits(n int) bool {
	return isMask(n ^ (n >> 1))
}

func isMask(mask int) bool {
	nextValueOfMask := mask + 1
	return isPowOfTwo(nextValueOfMask)
}

func isPowOfTwo(n int) bool {
	return getValueOfLowestBit(n) == n
}

func getValueOfLowestBit(n int) int {
	return n & (-n)
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-number-with-alternating-bits/submissions/
*/
