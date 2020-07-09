package 链表

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

/*
	题目链接: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/submissions/
	总结:
		1. 这题的解题思路比较Geek, 想法就是把2个链表串起来..2个指针将该链表走一遍，就能消除长度差了。
*/
