package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxDepth int // 记录树的最深深度
var sum int      // 记录树最深深度的节点和
func deepestLeavesSum(root *TreeNode) int {
	maxDepth = 0
	deepestLeavesSumExec(root, 0)
	return sum
}
func deepestLeavesSumExec(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	// 这里采取先、中、后遍历都可以
	if depth > maxDepth {
		maxDepth = depth
		sum = 0
	}
	if depth == maxDepth {
		sum += root.Val
	}
	deepestLeavesSumExec(root.Left, depth+1)
	deepestLeavesSumExec(root.Right, depth+1)
}
/*
	总结
	1. 这题考查的是二叉树的遍历。
	2. 我使用2个外部变量分别记录当前最深的深度及最深深度的节点和 + 二叉树DFS AC了这个题。

*/
