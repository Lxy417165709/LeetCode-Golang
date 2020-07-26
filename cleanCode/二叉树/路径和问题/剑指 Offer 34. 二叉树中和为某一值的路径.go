package 路径和问题

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	return getPaths(root, sum, []int{})
}

func getPaths(root *TreeNode, remainSum int, curPath []int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil && remainSum == root.Val {
		return [][]int{NewSlice(append(curPath, root.Val))}
	}
	paths := make([][]int, 0)
	paths = append(paths, getPaths(root.Left, remainSum-root.Val, append(curPath, root.Val))...)
	paths = append(paths, getPaths(root.Right, remainSum-root.Val, append(curPath, root.Val))...)
	return paths
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
*/


