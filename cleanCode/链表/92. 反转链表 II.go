package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preNodeOfReversedListHead := getKthNode(dummyHead, m)
	reversedListTail := preNodeOfReversedListHead.Next
	reversedListHead, remainListHead := reverse(preNodeOfReversedListHead.Next, n-m+1)
	preNodeOfReversedListHead.Next = reversedListHead
	reversedListTail.Next = remainListHead
	return dummyHead.Next
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

func reverse(head *ListNode, times int) (*ListNode, *ListNode) {
	var pre, next *ListNode
	cur := head
	for i := 1; i <= times; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre, cur
}

/*
	题目链接:
	总结:
		1. 题目要求一次遍历实现翻转，所以上面的代码简洁性可能不是很好。
*/
