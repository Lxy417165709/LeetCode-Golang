package 链表

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 初始化。
	curA, curB := headA, headB

	// 2. 循环走，直到指针相同。 (以下走法可用于消除长度差)
	for curA != curB {
		// 2.1 如果A指针为空，此时指向B头，否则指向下一个节点。
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		// 2.2 如果B指针为空，此时指向A头，否则指向下一个节点。
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	// 3. 返回。
	return curA
}
