package main

// 原地反转字符串
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/reverse-string/	翻转字符串
*/

/*
	总结
	1. 这题还可以拓展为翻转字符串中的一部分。
*/
