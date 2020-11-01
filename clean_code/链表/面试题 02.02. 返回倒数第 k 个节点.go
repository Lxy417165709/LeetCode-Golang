package 链表

func kthToLast(head *ListNode, k int) int {
	return getLastKthNode(head, k).Val
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

/*
	题目链接: https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
*/
