package 链表


// todo: 有更好的写法~ 可以统一。

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

		// 2.2 如果不存在重复节点，则递进，否则删除重复节点后递进。
		if pre.Next == cur {
			pre = cur
		} else {
			pre.Next = cur
			pre = cur
		}
	}

	// 3. 返回。
	return dummyHead.Next
}
