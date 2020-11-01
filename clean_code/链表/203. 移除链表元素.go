package 链表



func removeElements(head *ListNode, val int) *ListNode {
	return deleteNode(head,val)
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

/*
	题目链接: https://leetcode-cn.com/problems/remove-linked-list-elements/submissions/
*/
