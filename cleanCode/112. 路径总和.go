package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if isNil(root) {
		return false
	}
	nextSearchSum := sum - root.Val
	if isLeaf(root) {
		return nextSearchSum == 0
	}
	return hasPathSum(root.Left, nextSearchSum) || hasPathSum(root.Right, nextSearchSum)
}

func isNil(root *TreeNode) bool {
	return root == nil
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}



/*
	题目链接: https://leetcode-cn.com/problems/path-sum/
*/
