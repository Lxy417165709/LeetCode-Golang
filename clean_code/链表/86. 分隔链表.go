package 链表


// ------------------- 方法1: 时空复杂度 O(n),O(1) -------------------
func partition(head *ListNode, x int) *ListNode {
	lessPart, greaterOrEqualPart := parseToTwoPart(head, x)
	return link(lessPart, greaterOrEqualPart)
}

func parseToTwoPart(list *ListNode, ref int) (*ListNode, *ListNode) {
	dummyHeadOfLessPart := &ListNode{}
	dummyHeadOfGreaterOrEqualPart := &ListNode{}
	curL := dummyHeadOfLessPart
	curGE := dummyHeadOfGreaterOrEqualPart
	cur := list
	for cur != nil {
		if cur.Val < ref {
			curL.Next = cur
			curL = curL.Next
		} else {
			curGE.Next = cur
			curGE = curGE.Next
		}
		cur = cur.Next
	}
	curL.Next, curGE.Next = nil, nil
	return dummyHeadOfLessPart.Next, dummyHeadOfGreaterOrEqualPart.Next
}

func link(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	getTail(listA).Next = listB
	return listA
}

func getTail(notNilList *ListNode) *ListNode {
	cur := notNilList
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

/*
	题目链接: https://leetcode-cn.com/problems/partition-list/
*/


