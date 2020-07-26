package Geek

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	return len(s1)==len(s2) && strings.Contains(s2+s2,s1)
}

/*
	题目链接: https://leetcode-cn.com/problems/string-rotation-lcci/
	总结:
		1. 判断A是B的子串时，应该可以使用KMP算法。
*/
