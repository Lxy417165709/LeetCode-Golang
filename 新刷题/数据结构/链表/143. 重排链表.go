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

// spiltList 根据序号的奇偶性拆分链表，返回具有伪头的奇偶序号链表。
func spiltList(head *ListNode) (*ListNode, *ListNode) {
	oddListDummyHead := &ListNode{}
	evenListDummyHead := &ListNode{}

	num := 1
	cur := head
	curOfOddList := oddListDummyHead
	curOfEvenList := evenListDummyHead
	for cur != nil {
		if num&1 == 1 {
			curOfOddList.Next = cur
			curOfOddList = curOfOddList.Next
		} else {
			curOfEvenList.Next = cur
			curOfEvenList = curOfEvenList.Next
		}
		cur = cur.Next
		num++
	}
	curOfOddList.Next = nil
	curOfEvenList.Next = nil
	return oddListDummyHead, evenListDummyHead
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
