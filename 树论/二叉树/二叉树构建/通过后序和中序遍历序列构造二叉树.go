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

func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeExec(inorder, postorder)
}

func buildTreeExec(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	rootVal := postorder[n-1]
	index := 0
	for i := 0; i < n; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTreeExec(inorder[:index], postorder[:index]),
		Right: buildTreeExec(inorder[index+1:], postorder[index:n-1]),
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/		从中序与后序遍历序列构造二叉树
*/

/*
	总结
	1. 从后序和中序遍历序列构建二叉树的过程:
						(1) 先从后序遍历中找到根节点。
								(最后一个元素就是根节点)
						(2) 从中序遍历序列中找到根节点对应的索引，记为index，将中序遍历序列分为2部分，则每部分为左右子树的中序遍历序列。
								(inorder[:index]为左子树中序遍历序列，inorder[index+1:]为右子树中序遍历序列)
						(3) 从后序遍历序列中获取左右子树的后序遍历序列。
								(postorder[:index]为左子树后序遍历序列，postorder[index:n-1]为右子树后序遍历序列)
						(4) 递归构建左右子树
*/
