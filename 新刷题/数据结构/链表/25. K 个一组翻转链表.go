package 链表

func reverseKGroup(head *ListNode, k int) *ListNode {
	return reverse(head, k)
}

func getLength(head *ListNode) int {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}

// reverseByLoop k个一组翻转。 (循环)
func reverseByLoop(head *ListNode, k int) *ListNode {
	// 1. 获取链表长度。
	length := getLength(head)

	// 2. 循环翻转。
	dummyHead := &ListNode{
		Next: head,
	}
	pre, cur, next := dummyHead, dummyHead.Next,
		(*ListNode)(nil)
	for i := 1; i <= length/k; i++ {
		// 2.1 让 pre 指向该组链表的第t+1个节点。
		for t := 1; t <= k-1; t++ {
			next = cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}

		// 2.2 进入下一组链表。
		pre = cur
		cur = pre.Next
	}

	// 3. 返回。
	return dummyHead.Next
}

// reverse k个一组翻转。
func reverse(head *ListNode, k int) *ListNode {
	// 1. 获取链表长度。
	length := getLength(head)

	// 2. 递归出口。
	if length <= 1 {
		return head
	}
	if length < k {
		return head
	}

	// 3. 翻转前k个。
	pre, cur, next := (*ListNode)(nil), head, (*ListNode)(nil)
	for i := 1; i <= k; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 4. 递归。
	head.Next = reverse(cur, k)

	// 5. 返回。
	return pre
}
