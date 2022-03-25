package 链表

// removeNthFromEnd 删除链表倒数第n个节点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. 初始化。
	dummyHead := &ListNode{
		Next: head,
	}

	// 2. 寻找倒数第n+1个节点，记录到 first 指针中。
	last := dummyHead
	first := dummyHead
	for i := 1; i <= n; i++ {
		last = last.Next
	}
	for last.Next != nil {
		last = last.Next
		first = first.Next
	}

	// 3. 删除倒数第n个节点。
	first.Next = first.Next.Next

	// 4. 返回。
	return dummyHead.Next
}
