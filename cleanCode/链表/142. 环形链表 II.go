package 链表

// -------------------- 方法1: 时空复杂度 O(n),O(1) --------------------
func detectCycle(head *ListNode) *ListNode {
	return getEntranceNodeOfCycle(head)
}

func getEntranceNodeOfCycle(head *ListNode) *ListNode {
	nodeInCycle := getNodeInCycle(head)
	if nodeInCycle == nil {
		return nil
	}
	curC, curH := nodeInCycle, head
	for curC != curH {
		curC = curC.Next
		curH = curH.Next
	}
	return curC
}

func getNodeInCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

/*
	题目链接: https://leetcode-cn.com/problems/linked-list-cycle-ii/
*/
