package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 递归 + 外部变量 实现后序遍历
var postorderSequence []int
func postorderTraversal(root *TreeNode) []int {
	postorderSequence = make([]int, 0)
	postorderTraversalExec(root)
	return postorderSequence
}

func postorderTraversalExec(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversalExec(root.Left)
	postorderTraversalExec(root.Right)
	postorderSequence = append(postorderSequence, root.Val)
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/submissions/
*/