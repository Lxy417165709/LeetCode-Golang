package 链表

// detectCycle 获取链表环入口。
func detectCycle(head *ListNode) *ListNode {
	// 1. 判断是否有环，无环返回空。
	if !hasCycle(head) {
		return nil
	}

	// 2. 获取环内节点。
	// 证明:
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 3. 获取环入口。
	// slow、fast 相遇时，slow 距离链表环相差 a 步。 (a为链表非环节点数)
	// 因此，让 cur 指向表头，当 cur == slow，此时 slow 走了 a 步，即到达了链表环入口。
	// 证明: https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
	cur := head
	for cur != slow {
		cur = cur.Next
		slow = slow.Next
	}

	// 4. 返回。
	return cur
}
