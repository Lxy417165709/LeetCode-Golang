package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const INF = 10000000000

func isCousins(root *TreeNode, x int, y int) bool {
	return isBrother(root, x, y) && !isFatherSame(root, x, y)
}

func isBrother(root *TreeNode, x, y int) bool {
	return getMinLay(root, 0, x) == getMinLay(root, 0, y)
}

func isFatherSame(root *TreeNode, x, y int) bool {
	return getFatherInMinLay(root, x) == getFatherInMinLay(root, y)
}

func getFatherInMinLay(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left.Val == value {
		return root
	}
	if root.Right != nil && root.Right.Val == value {
		return root
	}
	leftMinFather, rightMinFather := getFatherInMinLay(root.Left, value), getFatherInMinLay(root.Right, value)
	if leftMinFather != nil {
		return leftMinFather
	}
	if rightMinFather != nil {
		return rightMinFather
	}
	return nil
}

func getMinLay(root *TreeNode, nowLay int, value int) int {
	if root == nil {
		return INF
	}
	if root.Val == value {
		return nowLay
	}
	return min(getMinLay(root.Left, nowLay+1, value), getMinLay(root.Right, nowLay+1, value))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/cousins-in-binary-tree/
*/
