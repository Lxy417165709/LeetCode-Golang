package main

func firstUniqChar(s string) int {
	count := make(map[uint8]int)
	for i:=0;i<len(s);i++{
		count[s[i]]++
	}
	for i:=0;i<len(s);i++{
		if count[s[i]]==1{
			return i
		}
	}
	return -1
}
/*
	题目链接:
		https://leetcode-cn.com/problems/first-unique-character-in-a-string/		字符串中的第一个唯一字符
*/
/*
	总结
	1. 这题采用哈希的方式时，时间复杂度为O(n)，空间复杂度为O(n)
	2. 当然，这题也有时间复杂度O(n^2)，空间O(1)的算法，就是使用暴力的方式，对每一个字符都向后找是否有重复的，
		有的话则在对应位置上用其他符号代替，防止重复扫描，没有的话就说明该字符是唯一字符。
*/