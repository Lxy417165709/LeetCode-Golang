package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// reversePrint 反向打印链表。
func reversePrint(head *ListNode) []int {
	return scanList(reverseList(head))
}

// scanList 遍历链表。
func scanList(listHead *ListNode) []int {
	result := make([]int, 0)
	cur := listHead
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

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
