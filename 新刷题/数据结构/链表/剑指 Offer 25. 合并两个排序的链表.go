package 链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeSortedLists(l1, l2)
}

// mergeSortedLists 将两个有序的链表合并为一个有序的链表。
func mergeSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 空判断。
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 2. 递归。
	if l1.Val == l2.Val {
		l1.Next = mergeSortedLists(l1.Next, l2)
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = mergeSortedLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeSortedLists(l1.Next, l2)
		return l1
	}
}
