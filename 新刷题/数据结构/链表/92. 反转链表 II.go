package 链表

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 1. 初始化。
	dummyHead := &ListNode{
		Next: head,
	}
	pre := dummyHead

	// 2. 让 pre 走到翻转链表的前一个节点。
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	// 3. 执行翻转。
	cur := pre.Next
	for i := 1; i <= right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	// 4. 返回。
	return dummyHead.Next
}
