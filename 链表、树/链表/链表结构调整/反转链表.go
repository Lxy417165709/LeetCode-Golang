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

// 递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {

}

/*
	题目链接: https://leetcode-cn.com/problems/reverse-linked-list/
*/
