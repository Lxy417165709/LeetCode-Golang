package main


func sortList(head *ListNode) *ListNode {
	listLength := getLength(head)
	for cutSize := 1; cutSize < listLength; cutSize *= 2 {
		head = partSort(head,cutSize)
	}
	return head
}

// 这一部分急需重构！
func partSort(list *ListNode,sortLength int) *ListNode {
	dummyHead := &ListNode{Next:list}
	preCur,cur := dummyHead,dummyHead.Next
	for curNodeOfSortedList != nil {
		leftSortedList := cur
		rightSortedList := cut(leftSortedList, sortLength)
		cur = cut(rightSortedList, sortLength)


		preCur.Next = mergeSortedLists(leftSortedList, rightSortedList)
		preCur = getTailNodeOfNotNilList(preCur)
		preCur.Next = cur
	}
	return dummyHead.Next
}
func link(listA*ListNode,listB *ListNode) {
	tailNodeOfListA := getTailNodeOfNotNilList(listA)
	tailNodeOfListA.Next = listB
}
func getLinkList(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	tailNodeOfListA := getTailNodeOfNotNilList(listA)
	tailNodeOfListA.Next = listB
	return listA
}



func getTailNodeOfNotNilList(list *ListNode) *ListNode {
	cur := list
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func getLength(list *ListNode) int {
	cur := list
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}

func cut(list *ListNode, cutSize int) *ListNode {
	cutListDummyHead := &ListNode{}
	curNodeOfCutList := cutListDummyHead
	curNodeOfList := list
	for i := 1; i <= cutSize && curNodeOfList != nil; i++ {
		curNodeOfCutList.Next = curNodeOfList
		curNodeOfList = curNodeOfList.Next
		curNodeOfCutList = curNodeOfCutList.Next
	}
	nextNodeOfTail := curNodeOfCutList.Next
	curNodeOfCutList.Next = nil
	return nextNodeOfTail
}

func mergeSortedLists(listA *ListNode, listB *ListNode) *ListNode {
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
	mergedListHead.Next = getNotNilList(listA, listB)
	return dummyMergedListHead.Next
}

func getNotNilList(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	panic("不存在非nil链表")
}
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	题目链接: https://leetcode-cn.com/problems/sort-list/
	总结
		1. 这题搁置了好久没做了。
		2. 这题需要用到归并排序，而且不能使用递归的归并排序，只能使用迭代的，这样才能保证O(1)时间复杂度。
*/
