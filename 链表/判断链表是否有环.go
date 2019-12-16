package main

// 哈希解法	(判断是否走了回头路)
func hasCycle(head *ListNode) bool {
	isVisited := make(map[*ListNode]int)
	for head != nil {
		if _, ok := isVisited[head]; ok {
			return true
		}
		isVisited[head] = 1
		head = head.Next
	}
	return false
}

// 双指针解法 (慢指针每次走一步，快指针每次走两步，相遇则有环)
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
	题目链接:
		https://leetcode-cn.com/problems/linked-list-cycle/		环形链表
*/
