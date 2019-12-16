package main


var lastSum int
func convertBST(root *TreeNode) *TreeNode {
	lastSum = 0
	convertBSTExec(root)
	return root
}

func convertBSTExec(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTExec(root.Right)
	root.Val += lastSum
	lastSum = root.Val
	convertBSTExec(root.Left)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/convert-bst-to-greater-tree/comments/		把二叉搜索树转换为累加树
*/

/*
	总结
	1. 这题就是反序的中序遍历
*/