package 链表

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/struct_util"

// mergeKLists 合并 K 个有序的链表。
func mergeKLists(lists []*ListNode) *ListNode {
	// 1. 建小顶堆。
	smallTopHeap := struct_util.NewMyHeap(len(lists), func(a, b interface{}) bool {
		return a.(*ListNode).Val < b.(*ListNode).Val
	})

	// 2. 压入链表头，初始化堆。
	for _, listHead := range lists {
		if listHead == nil {
			continue
		}
		smallTopHeap.Push(listHead)
	}

	// 3. 利用堆，合并链表。
	dummyHeadOfMergeList := &ListNode{}
	for curHead := dummyHeadOfMergeList; smallTopHeap.Size() != 0; curHead = curHead.Next {
		// 3.1 获取堆顶元素，推入下一元素。
		curSmallestNode := smallTopHeap.Pop().(*ListNode)
		if curSmallestNode.Next != nil {
			smallTopHeap.Push(curSmallestNode.Next)
		}

		// 3.2 构建链表结构。
		curHead.Next = curSmallestNode
	}

	// 4. 返回。
	return dummyHeadOfMergeList.Next
}
