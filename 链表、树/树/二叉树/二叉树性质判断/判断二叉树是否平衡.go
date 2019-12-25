package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 判断二叉树是否为平衡二叉树		(写法1)
func isBalanced(root *TreeNode) bool {
	ans, _ := isBalancedExec(root)
	return ans
}

// 第一个返回值是该树是否平衡，第二个是该树的高度
func isBalancedExec(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftIsBalanced, leftHeight := isBalancedExec(root.Left)
	rightIsBalanced, rightHeight := isBalancedExec(root.Right)
	return leftIsBalanced && rightIsBalanced && abs(leftHeight-rightHeight) <= 1, max(leftHeight, rightHeight) + 1
}

// 判断二叉树是否为平衡二叉树		(写法2)
func isBalanced(root *TreeNode) bool {
	ans := isBalancedExec(root)
	return ans != -1
}

// 返回树高 (-1表示不是平衡二叉树)
func isBalancedExec(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := isBalancedExec(root.Left)
	rightHeight := isBalancedExec(root.Right)
	if leftHeight != -1 && rightHeight != -1 && abs(leftHeight-rightHeight) <= 1 {
		return max(leftHeight, rightHeight) + 1
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接:
*/

/*
	总结
	1. 写法1和写法2的区别在于，写法2只有一个返回值，返回树高，通过树高来判断是否为平衡二叉树(如果是-1就说不是平衡二叉树)。
*/
