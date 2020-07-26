package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		removeNodesWhoseValIsEqualToHead(cur)
		cur = cur.Next
	}
	return head
}
func removeNodesWhoseValIsEqualToHead(head *ListNode) {
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == head.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/remove-duplicate-node-lcci/comments/
	总结
		1. 以上的解法采用暴力搜索，时空复杂度为 O(n^2) O(1)
		2. 这题也可以采用哈希解决，时空复杂度为 O(n) O(n)
*/
