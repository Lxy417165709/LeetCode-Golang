package 链表

// reorderList 重排链表。
func reorderList(head *ListNode) {
	// 1. 获取链表左中点。
	midNode := getMidNode(head)

	// 2. 拆分链表。
	list1, list2 := head, reverse(midNode.Next)
	midNode.Next = nil // 断开 list1 与 list2 的连接。

	// 3. 合并链表。
	mergeList(list1, list2)
}

// getMidNode 获取中间节点，链表长度为偶数时，返回序号为偶数的中点。 (序号从1开始)
func getMidNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// reverse 翻转链表。
func reverse(head *ListNode) *ListNode {
	pre, cur, next := (*ListNode)(nil), head, (*ListNode)(nil)
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// mergeList 合并链表。
func mergeList(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	list1.Next = mergeList(list2, list1.Next)
	return list1
}

