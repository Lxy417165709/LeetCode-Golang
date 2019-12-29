package main

// 一般解法，使用哑头节点
func oddEvenList(head *ListNode) *ListNode {
	oddDummyHead, evenDummyHead := &ListNode{-1, nil}, &ListNode{-1, nil}
	odd, even := oddDummyHead, evenDummyHead
	order := 1
	for head != nil {
		if order%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		order++
	}
	odd.Next = evenDummyHead.Next
	even.Next = nil // 记得清空
	return oddDummyHead.Next
}

// 不使用哑头节点 (官方解法)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// odd为奇链表表头, even为偶链表表头
	odd, even, tmp := head, head.Next, head.Next
	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = tmp
	return head
}

/*
	题目链接:
		https://leetcode-cn.com/problems/odd-even-linked-list/
*/
