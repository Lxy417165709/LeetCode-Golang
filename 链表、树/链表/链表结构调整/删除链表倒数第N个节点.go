package main

// 双指针法 删除链表的倒数第N个节点，时空复杂度 O(n),O(1)
// 需要判断n合法性，非法情况如: n小于0 or n大于链表长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	// 定义哑头节点，这样可以方便删除头节点时的处理
	dummyHead := &ListNode{-1, head}

	left, right := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	// 此时left指向了要删除节点的父节点
	left.Next = left.Next.Next
	return dummyHead.Next
}

/*
	题目链接: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/comments/
*/

/*
	总结
	1. 思路是:
			(1) 将左右指针指向哑头节点
			(2) 右指针先走n步，之后左右指针一起走，直到右指针到达尾节点
			(3) 此时左指针指向了删除节点的前驱节点
			(4) 执行: left.Next = left.Next.Next，删除该节点
*/