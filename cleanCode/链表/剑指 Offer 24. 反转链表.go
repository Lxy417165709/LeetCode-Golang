package 链表

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*
	题目链接: https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
*/
