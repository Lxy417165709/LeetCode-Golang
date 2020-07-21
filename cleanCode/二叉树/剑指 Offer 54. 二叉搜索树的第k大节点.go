package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const flagOfCannotFind = 100000000000000 // 该值不能存在于树中
var curOrder int

func kthLargest(root *TreeNode, k int) int {
	curOrder = 0
	return getKthLargest(root, k)
}

func getKthLargest(root *TreeNode, k int) int {
	if root == nil {
		return flagOfCannotFind
	}
	if rightResult := getKthLargest(root.Right, k); rightResult != flagOfCannotFind {
		return rightResult
	}
	if curOrder++; curOrder == k {
		return root.Val
	}
	if leftResult := getKthLargest(root.Left, k); leftResult != flagOfCannotFind {
		return leftResult
	}
	return flagOfCannotFind
}

/*
	题目链接: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
*/
