package 链表



func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

/*
	题目链接: https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
*/
