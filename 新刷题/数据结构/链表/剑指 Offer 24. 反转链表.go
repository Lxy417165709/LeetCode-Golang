package 链表

// reverseList 反转链表。
func reverseList(listHead *ListNode) *ListNode {
	pre, cur, next := (*ListNode)(nil), listHead, (*ListNode)(nil)
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
