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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeExec(preorder, inorder)
}

func buildTreeExec(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTreeExec(preorder[1:index+1], inorder[:index]),
		Right: buildTreeExec(preorder[index+1:], inorder[index+1:]),
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/submissions/		从前序与中序遍历序列构造二叉树
*/

/*
	总结
	1. 从先序和中序遍历序列构建二叉树的过程:
						(1) 先从先序遍历中找到根节点。
								(第一个元素就是根节点)
						(2) 从中序遍历序列中找到根节点对应的索引，记为index，将中序遍历序列分为2部分，则每部分为左右子树的中序遍历序列。
								(inorder[:index]为左子树中序遍历序列，inorder[index+1:]为右子树中序遍历序列)
						(3) 从先序遍历序列中获取左右子树的先序遍历序列。
								(preorder[1:index+1]为左子树先序遍历序列，preorder[index+1:]为右子树先序遍历序列)
						(4) 递归构建左右子树
*/
