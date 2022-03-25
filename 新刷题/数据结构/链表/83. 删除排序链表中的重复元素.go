package 链表

// deleteDuplicates 删除链表重复节点。 (重复的保留一个)
func deleteDuplicates(head *ListNode) *ListNode {
	// 1. 初始化。
	dummyHead := &ListNode{
		Next: head,
	}
	pre := dummyHead

	// 2. 去除重复节点。
	for cur := pre.Next; cur != nil; cur = pre.Next {
		// 2.1 标识重复节点。
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		// 2.2 如果存在重复节点则删除重复节点。
		if pre.Next != cur {
			pre.Next = cur
		}

		// 2.3 递进。
		pre = cur
	}

	// 3. 返回。
	return dummyHead.Next
}
