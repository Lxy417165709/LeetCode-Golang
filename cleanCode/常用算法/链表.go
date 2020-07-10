package main

// -------------------------- 获取 ----------------------------
func getEntranceNodeOfCycle(head *ListNode) *ListNode {
	nodeInCycle := getNodeInCycle(head)
	if nodeInCycle == nil {
		return nil
	}
	curC, curH := nodeInCycle, head
	for curC != curH {
		curC = curC.Next
		curH = curH.Next
	}
	return curC
}

func getNodeInCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
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

func getLastKthNode(head *ListNode, k int) *ListNode {
	aft := getKthNode(head, k)
	cur := head
	for aft.Next != nil {
		cur = cur.Next
		aft = aft.Next
	}
	return cur
}

func getKthNode(head *ListNode, k int) *ListNode {
	cur := head
	moveTimes := k - 1
	for cur != nil && moveTimes > 0 {
		moveTimes--
		cur = cur.Next
	}
	return cur
}

func getRightMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func getLeftMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func getTailNode(notNilList *ListNode) *ListNode {
	cur := notNilList
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
// ---------------------------- 判断 ----------------------------
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	leftMiddleNode := getLeftMiddleNode(head)
	rightPartList := reverseList(leftMiddleNode.Next)
	return isPrefix(head, rightPartList)
}

func isPrefix(patternList *ListNode, sampleList *ListNode) bool {
	pCur := patternList
	sCur := sampleList
	for pCur != nil && sCur != nil {
		if pCur.Val != sCur.Val {
			return false
		}
		pCur = pCur.Next
		sCur = sCur.Next
	}
	return sCur == nil
}

//  ---------------------------- 修改链表内部结构 ----------------------------
func reverseList(list *ListNode) *ListNode {
	var next, pre *ListNode
	cur := list
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func link(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	tailNodeOfListA := getTailNode(listA)
	tailNodeOfListA.Next = listB
	return listA
}
