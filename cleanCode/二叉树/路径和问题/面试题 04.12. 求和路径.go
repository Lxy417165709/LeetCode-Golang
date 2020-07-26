package 路径和问题

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return getCountOfPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func getCountOfPath(root *TreeNode, remainSum int) int {
	if root == nil {
		return 0
	}
	countOfPath := 0
	if remainSum == root.Val {
		countOfPath++
	}
	countOfPath += getCountOfPath(root.Left, remainSum-root.Val)
	countOfPath += getCountOfPath(root.Right, remainSum-root.Val)
	return countOfPath
}

// ---------------- 前缀和 + 回溯 -----------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var countOfPath map[int]int

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	countOfPath = make(map[int]int)
	countOfPath[0] = 1
	return getCountOfPath(root, 0, sum)
}

func getCountOfPath(root *TreeNode, curSum int, targetSum int) int {
	if root == nil {
		return 0
	}
	curSum += root.Val
	count := countOfPath[curSum-targetSum]
	countOfPath[curSum]++
	count += getCountOfPath(root.Left, curSum, targetSum)
	count += getCountOfPath(root.Right, curSum, targetSum)
	countOfPath[curSum]--
	return count
}

/*
	题目链接: https://leetcode-cn.com/problems/paths-with-sum-lcci/
	总结:
		1. 官方有使用前缀和AC这题的。
*/
