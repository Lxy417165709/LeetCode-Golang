package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 优美的解法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var next *ListNode = nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {

}
/*
	题目链接: https://leetcode-cn.com/problems/reverse-linked-list/
*/

