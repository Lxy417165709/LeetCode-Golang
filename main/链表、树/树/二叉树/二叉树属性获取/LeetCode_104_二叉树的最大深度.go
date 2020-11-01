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

// 求二叉树深度 (最大深度)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	return max(left, right) + 1
}

// 求二叉树深度 (最小深度)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 && right == 0 {
		return 1
	}
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	return min(left, right) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/		二叉树最小深度
		https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/		二叉树最大深度
*/