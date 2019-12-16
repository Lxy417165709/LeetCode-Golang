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

// 判断二叉树是否为单值二叉树	(写法1)
func isUnivalTree(root *TreeNode) bool {
	return isUnivalTreeExec(root)
}

func isUnivalTreeExec(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}
	return isUnivalTreeExec(root.Left) && isUnivalTreeExec(root.Right)
}

// 判断二叉树是否为单值二叉树	(写法2)
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeExec(root, root.Val)
}

func isUnivalTreeExec(root *TreeNode, number int) bool {
	if root == nil {
		return true
	}
	return root.Val == number && isUnivalTreeExec(root.Left, number) && isUnivalTreeExec(root.Right, number)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/univalued-binary-tree/submissions/		单值二叉树
*/
