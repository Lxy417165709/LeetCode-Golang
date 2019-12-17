package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针法获取环入口
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// 判断链表是否有环
	loopFlag := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			loopFlag = true
			break
		}
	}
	// 表示链表无环
	if loopFlag == false {
		return nil
	}
	// 表示链表有环
	loopEntrance := head
	for loopEntrance != slow {
		loopEntrance = loopEntrance.Next
		slow = slow.Next
	}
	return loopEntrance
}

/*
	题目链接:
		https://leetcode-cn.com/problems/linked-list-cycle-ii/		环形链表2
*/