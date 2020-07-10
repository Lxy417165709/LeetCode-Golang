package 链表

// ------------------------- 方法1: 暴力K路归并 -------------------------
// 执行用时：320 ms, 在所有 Go 提交中击败了 10.75% 的用户
// 内存消耗：5.3 MB, 在所有 Go 提交中击败了 100.00% 的用户
const INF = 1000000000
const cantFindOutIndexFlag = -1
func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for {
		minIndex := getIndexOfMinNode(lists)
		if minIndex == cantFindOutIndexFlag {
			break
		}
		cur.Next = lists[minIndex]
		cur = cur.Next
		lists[minIndex] = lists[minIndex].Next
	}
	return dummyHead.Next
}

func getIndexOfMinNode(lists []*ListNode) int {
	minIndex := cantFindOutIndexFlag
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		if minIndex == cantFindOutIndexFlag || lists[i].Val < lists[minIndex].Val {
			minIndex = i
		}
	}
	return minIndex
}

// ------------------------- 方法2: 2路归并实现K路归并 -------------------------
// 执行用时：4 ms, 在所有 Go 提交中击败了 99.41% 的用户
// 内存消耗：5.3 MB, 在所有 Go 提交中击败了 77.78% 的用户
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

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
	题目链接: https://leetcode-cn.com/problems/merge-k-sorted-lists/
	总结
		1. 方法1 K路归并找最小节点的时候 (即 getIndexOfMinNode 函数)，可以使用优先队列优化！
*/
