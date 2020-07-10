package main

// -------------- 时空复杂度: O(n),O(1) -------------------
// 执行用时：8 ms, 在所有 Go 提交中击败了99.17% 的用户
// 内存消耗：5.4 MB, 在所有 Go 提交中击败了100.00% 的用户
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	leftPart, rightPart := getTwoParts(head)
	reversedRightPart := reverse(rightPart)
	crossMerge(leftPart, reversedRightPart)
}

func getTwoParts(list *ListNode) (leftPart, rightPart *ListNode) {
	leftMiddleNode := getLeftMiddleNode(list)
	leftPart = list
	rightPart = leftMiddleNode.Next
	leftMiddleNode.Next = nil
	return
}

func reverse(list *ListNode) *ListNode {
	var pre, next *ListNode
	cur := list
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 要求listA的长度比listB大1
func crossMerge(listA, listB *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for listB != nil {
		nextA, nextB := listA.Next, listB.Next
		cur.Next = listA
		cur.Next.Next = listB
		cur = cur.Next.Next
		listA = nextA
		listB = nextB
	}
	cur.Next = listA
	return dummyHead.Next
}

// 例子: 1 -> 2 -> 3 -> 4，返回的节点是 2，而不是 3
// 例子: 1 -> 2 -> 3，     返回的节点是 2
func getLeftMiddleNode(list *ListNode) *ListNode {
	slow, fast := list, list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/*
	题目链接:
	总结:
		1. 只要熟悉流程，AC就很简单了。
		2. 结构清晰能让我们易于理解，易于理解就容易记住。
*/
