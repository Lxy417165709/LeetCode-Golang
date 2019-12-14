package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转二叉树 (先序)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 翻转二叉树 (中序)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	// 由于左右子树已经交换了，所以此时的左子树是原来的右子树
	invertTree(root.Left)
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/invert-binary-tree/submissions/		翻转二叉树
*/

/*
	总结
	1. 翻转二叉树可以使用先序遍历、中序遍历、后序遍历、层序遍历实现。
       如果采用中序遍历的话，请参考上面的代码。
*/
