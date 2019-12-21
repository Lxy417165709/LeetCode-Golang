package main

// 优美的解法 (根本在于消除长度差)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

func main() {

}

/*
	题目链接: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
*/