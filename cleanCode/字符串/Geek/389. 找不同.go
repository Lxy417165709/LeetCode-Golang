package Geek

func findTheDifference(s string, t string) byte {
	return byte(getSum(t) - getSum(s))
}

func getSum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}

/*
	题目链接: https://leetcode-cn.com/problems/find-the-difference/
	总结:
		1. 这题比较geek的就是，不使用哈希，而是记录字符和，最后相减就能得到不同的字符了。
		2. 也可以采用位运算，将所有s、t的字符异或就可以得到最终结果了。
*/
