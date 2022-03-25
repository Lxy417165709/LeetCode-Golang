package 链表

// sortList 排序链表。
// 空间复杂度: O(1)。
// 时间复杂度: O(nlogn)
func sortList(head *ListNode) *ListNode {
	// 1. 初始化。
	dummyHead := &ListNode{
		Next: head,
	}

	// 2. 获取链表长度。
	length := getLength(head)

	// 3. 由下到上归并排序。
	for size := 1; size <= length; size <<= 1 {
		tail := dummyHead
		cur := tail.Next
		for cur != nil {
			// 3.1 切分为3段链表。
			list1 := cur            // 第一段。
			list2 := cut(cur, size) // 第二段。
			cur = cut(list2, size)  // 第三段。

			// 3.2 合并。
			tail.Next = merge(list1, list2)

			// 3.3 维护，保证 tail 指向合并后链表的尾部。
			tail = getTailNode(tail)
		}
	}

	// 4. 返回。
	return dummyHead.Next
}

// getLength 获取链表长度。
func getLength(head *ListNode) int {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}

// cut 切分链表为 head[ :min(size, len(head))]，返回下一段链表的表头。
func cut(head *ListNode, size int) *ListNode {
	// 1. 让 cur 指向下一端链表的前一个节点。
	cur := head
	for i := 1; i <= size-1 && cur != nil; i++ {
		cur = cur.Next
	}

	// 2. 断开第一段、第二段链表的连接，返回第二段链表的表头。
	if cur != nil {
		nextListHead := cur.Next
		cur.Next = nil
		return nextListHead
	}

	// 3. 空返回。
	return nil
}

// merge 两个链表合并。
func merge(list1, list2 *ListNode) *ListNode {
	// 1. 初始化。
	dummyHead := &ListNode{}

	// 2. 合并。
	tail := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	}
	if list2 == nil {
		tail.Next = list1
	}

	// 3. 返回。
	return dummyHead.Next
}

// getTailNode 获取尾节点。
func getTailNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
