package main

// 原地反转字符串
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// 原地翻转bytes的某部分[first, last]
// first,last分别对应开始反转的起始索引、终止索引
func reverse(bytes []byte, first int, last int) {
	n := last - first + 1
	for i := 0; i < n/2; i++ {
		bytes[first+i], bytes[last-i] = bytes[last-i], bytes[first+i]
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/reverse-string/					翻转字符串
		https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/		反转字符串中的单词 III
*/

