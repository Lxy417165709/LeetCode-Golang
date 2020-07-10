package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddList, evenList := getOddAndEvenList(head)
	return link(oddList, evenList)
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
	题目链接: https://leetcode-cn.com/problems/odd-even-linked-list/
*/
