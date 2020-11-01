package 链表间关系判断

func isSame(listA, listB *ListNode) bool {
	for listA != nil && listB != nil {
		if listA.Val != listB.Val {
			return false
		}
		listA = listA.Next
		listB = listB.Next
	}
	return listA == listB
}
