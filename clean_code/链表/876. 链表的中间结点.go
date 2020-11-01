package 链表

// 链表长度是奇数时返回中间节点，偶数时返回右中间节点。
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 链表长度是奇数时返回中间节点，偶数时返回左中间节点。
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/*
	题目链接: https://leetcode-cn.com/problems/middle-of-the-linked-list/
*/
