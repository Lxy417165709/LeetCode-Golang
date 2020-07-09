package main

// -------------------------- 获取 ----------------------------
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
