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

// 获取完全二叉树的节点个数
// 获取完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	return countNodesExec(root)
}

func countNodesExec(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getCompleteBinaryTreeDepth(root.Left)
	rightDepth := getCompleteBinaryTreeDepth(root.Right)
	leftTreeNodes, rightTreeNodes := 0, 0

	if leftDepth == rightDepth {
		// leftDepth == rightDepth时，表明左子树是满二叉树
		leftTreeNodes = getFullBinaryTreeNodes(leftDepth)
		rightTreeNodes = countNodesExec(root.Right)
	} else {
		// leftDepth != rightDepth时，表明右子树是满二叉树
		leftTreeNodes = countNodesExec(root.Left)
		rightTreeNodes = getFullBinaryTreeNodes(rightDepth)
	}
	return leftTreeNodes + rightTreeNodes + 1
}

// 获取满二叉树的节点个数 (传入树的深度)
func getFullBinaryTreeNodes(depth int) int {
	return (1 << uint8(depth)) - 1
}

// 获取完全二叉树的高度 (每次都向左子树走)
func getCompleteBinaryTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getCompleteBinaryTreeDepth(root.Left) + 1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/count-complete-tree-nodes/			完全二叉树的节点个数
*/

/*
	总结
	1. 注意:
		完全二叉树			 -> 	左右子树是完全二叉树
		左右子树是完全二叉树 ≠> 	完全二叉树
	2. 想要获取完全二叉树的深度，只需要每次都走左边，尽头的深度就是完全二叉树的深度。
	3. 满二叉树的节点运算公式是: 2^depth - 1 == (1<<depth) - 1
*/
