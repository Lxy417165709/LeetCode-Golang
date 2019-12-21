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

// 求二叉树左叶子节点之和		(解法1)
func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesExec(root)
}

func sumOfLeftLeavesExec(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeavesExec(root.Right)
	}
	return sumOfLeftLeavesExec(root.Left) + sumOfLeftLeavesExec(root.Right)
}

// 求二叉树左叶子节点之和		(解法2)
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumOfLeftLeavesExec(root.Left, 0) + sumOfLeftLeavesExec(root.Right, 1)
}

func sumOfLeftLeavesExec(root *TreeNode, flag int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if flag == 0 {
			return root.Val
		}
		return 0
	}
	return sumOfLeftLeavesExec(root.Left, 0) + sumOfLeftLeavesExec(root.Right, 1)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/sum-of-left-leaves/			左叶子之和
*/

/*
	总结
	1. 解法1和解法2差不多，只是解法1判断左子树节点是从根出发，而解法2通过传入一个标志，
	   标识该树是左子树还是右子树。(这里我采用了0标识左子树，1标识右子树)
*/
