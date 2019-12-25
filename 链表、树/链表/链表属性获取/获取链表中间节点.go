package main

// 优美的解法 (使用快慢指针法，快指针每次走2步，慢指针每次走1步)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	//  slow,fast:=head,head		如果是这样写的话，如果链表节点为偶数，则返回偏右的中间节点
	//  slow,fast:=head,head.Next	如果是这样写的话，如果链表节点为偶数，则返回偏左的中间节点(注意判断head非空)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/*
	题目链接: https://leetcode-cn.com/problems/middle-of-the-linked-list/
*/
