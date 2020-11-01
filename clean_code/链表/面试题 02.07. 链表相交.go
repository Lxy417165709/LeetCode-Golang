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
	题目链接: https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/submissions/
	总结:
		1. 这题和 _160. 相交链表_ 是一样的
*/
