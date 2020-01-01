package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	return insertIntoMaxTreeExec(root, val)
}

func insertIntoMaxTreeExec(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val < val {
		return &TreeNode{val, root, nil}
	}

	root.Right = insertIntoMaxTreeExec(root.Right, val)
	return root

}
/*
	总结
	1. 这题很水...每次向右边走就可以了(如果值小于根节点)
	2. 关键在于理解题目。
*/
