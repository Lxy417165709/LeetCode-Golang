package main

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) | head.Val
		head = head.Next
	}
	return result
}
/*
	题目链接:
		https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/solution/wei-yun-suan-by-wf0312/    二进制链表转整数
*/
/*
	总结
	1. 这题目利用了进制转换的知识。
*/
