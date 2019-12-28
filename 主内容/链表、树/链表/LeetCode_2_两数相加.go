package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	cur := dummyHead
	p1, p2, carry := l1, l2, 0
	for p1 != nil || p2 != nil || carry != 0 {
		if p1 != nil {
			carry += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			carry += p2.Val
			p2 = p2.Next
		}
		cur.Next = &ListNode{carry % 10, nil}
		cur = cur.Next
		carry /= 10
	}
	return dummyHead.Next
}

/*
	题目链接:
		https://leetcode-cn.com/problems/add-two-numbers/submissions/		两数相加
*/
/*
	总结
	1. 这题的链表已经是逆序的，所以直接相加就可以了。如果不是逆序的，那么必须先把两个链表
		翻转，再执行相加操作，最后再把得出的结果链表翻转，那这就是答案了。
*/