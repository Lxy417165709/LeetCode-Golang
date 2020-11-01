package main

/*
	给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。

	返回移除了所有不包含 1 的子树的原二叉树。

	( 节点 X 的子树为 X 本身，以及所有 X 的后代。)
*/

// 移除所有不包含1的子树
func pruneTree(root *TreeNode) *TreeNode {
	return pruneTreeExec(root)
}

func pruneTreeExec(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTreeExec(root.Left)
	root.Right = pruneTreeExec(root.Right)
	if root.Left == nil && root.Right == nil && root.Val != 1 {
		return nil
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/binary-tree-pruning/		二叉树剪枝
*/

/*
	总结
	1. 这题还可以拓展喔，比如删除不包含某个子集元素的子树。 (这题中，子集为 [ 1 ])
*/