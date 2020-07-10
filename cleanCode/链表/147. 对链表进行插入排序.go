package 链表

const INF = 1000000000000

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -INF}
	cur := head
	for cur != nil {
		next := cur.Next
		insertIntoList(dummyHead, cur)
		cur = next
	}
	return dummyHead.Next
}

func insertIntoList(dummyHead *ListNode, node *ListNode) {
	// 这个函数会影响 node 的 Next指针。
	preNodeOfInserting := getPreNodeOfInserting(dummyHead, node)
	node.Next = preNodeOfInserting.Next
	preNodeOfInserting.Next = node
}

func getPreNodeOfInserting(dummyHead, node *ListNode) *ListNode {
	cur := dummyHead
	for !(cur.Val <= node.Val && (cur.Next == nil || node.Val <= cur.Next.Val)) {
		cur = cur.Next
	}
	return cur
}

/*
	题目链接: https://leetcode-cn.com/problems/insertion-sort-list/
	总结
		1. INF要定义大些，要超过INT32的最大值。
*/
