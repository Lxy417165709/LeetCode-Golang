package 链表

// ----------------------------- 迭代解法 -----------------------------
func swapPairs(head *ListNode) *ListNode {
	oddList, evenList := getOddAndEvenList(head)
	return crossMerge(evenList, oddList)
}

func getOddAndEvenList(list *ListNode) (*ListNode, *ListNode) {
	oddDummyHead, EvenDummyHead := &ListNode{}, &ListNode{}
	curO, curE := oddDummyHead, EvenDummyHead
	cur := list
	nowIndexOfNode := 1
	for cur != nil {
		if nowIndexOfNode%2 == 1 {
			curO.Next = cur
			curO = curO.Next
		} else {
			curE.Next = cur
			curE = curE.Next
		}
		cur = cur.Next
		nowIndexOfNode++
	}
	curO.Next, curE.Next = nil, nil
	return oddDummyHead.Next, EvenDummyHead.Next
}

func crossMerge(listA, listB *ListNode) *ListNode {
	// 要求listA较短
	dummyHead := &ListNode{}
	cur := dummyHead
	for listA != nil {
		nextA, nextB := listA.Next, listB.Next
		cur.Next = listA
		cur.Next.Next = listB
		cur = cur.Next.Next
		listA = nextA
		listB = nextB
	}
	cur.Next = listB
	return dummyHead.Next
}

// ----------------------------- 递归解法 -----------------------------
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

/*
	题目链接: https://leetcode-cn.com/problems/swap-nodes-in-pairs/submissions/
	总结:
		1. 这题还能用递归的。
		2. crossMerge函数与 _143. 重排链表_ 的很类似。
*/
