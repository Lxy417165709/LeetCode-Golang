package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 获取二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}


/*
	总结
	1. 二叉树的节点个数 == 左子树节点个数 + 右子树节点个数 + 1
*/
