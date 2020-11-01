package 链表

// 第一次写出的代码
func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}

// 看了题解后写出的代码
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

/*
	题目链接: https://leetcode-cn.com/problems/delete-middle-node-lcci/submissions/
	总结
		1. 这题比较Geek, 不删除点，而是通过替换值，来达到删除的目的
		2. 这题和 _237. 删除链表中的节点_ 是一样的。
*/
