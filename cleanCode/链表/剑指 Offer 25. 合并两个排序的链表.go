package 链表



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

/*
	题目链接: https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
	总结:
		1. 这题和 _21. 合并两个有序链表_ 是一样的。
*/
