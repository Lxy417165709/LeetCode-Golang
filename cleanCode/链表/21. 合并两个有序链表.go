package 链表

// --------------------- 迭代版 ---------------------
func mergeTwoLists(listA *ListNode, listB *ListNode) *ListNode {
	dummyMergedListHead := &ListNode{-1, nil}
	mergedListHead := dummyMergedListHead
	for listA != nil && listB != nil {
		if listA.Val > listB.Val {
			mergedListHead.Next = listB
			listB = listB.Next
		} else {
			mergedListHead.Next = listA
			listA = listA.Next
		}
		mergedListHead = mergedListHead.Next
	}
	if listA == nil {
		mergedListHead.Next = listB
	}
	if listB == nil {
		mergedListHead.Next = listA
	}
	return dummyMergedListHead.Next
}

// --------------------- 递归版 ---------------------
func mergeTwoLists(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	if listA.Val > listB.Val {
		listB.Next = mergeTwoLists(listA, listB.Next)
		return listB
	} else {
		listA.Next = mergeTwoLists(listB, listA.Next)
		return listA
	}
}
