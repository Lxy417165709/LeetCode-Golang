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

func bstFromPreorder(preorder []int) *TreeNode {
	return bstFromPreorderExec(preorder)
}

func bstFromPreorderExec(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := 0
	for index = 0; index < len(preorder); index++ {
		if preorder[index] > rootVal {
			break
		}
	}
	root := &TreeNode{
		Val:   rootVal,
		Left:  bstFromPreorderExec(preorder[1:index]),
		Right: bstFromPreorderExec(preorder[index:]),
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/submissions/		先序遍历构造二叉树
*/

/*
	总结
	1. 从先序构建二叉搜索树的过程:
			(1) 先从先序遍历序列中找到根节点。
					(第一个元素就是根节点)
			(2) 扫描先序遍历序列，找到一个index，使得preorder[index] > rootVal。
					(二叉搜索树的左子树都小于根节点的值，而右子树的都大于根节点的值)
			(3) 从先序遍历序列中获取左右子树的先序遍历序列。
					(preorder[1:index]为左子树先序遍历序列，preorder[index:]为右子树先序遍历序列)
			(4) 递归构建左右子树。
	2. 注意: 题目已经保证了先序遍历序列是合法的，能够构建出一颗二叉搜索树。
			( 比如: [8,5,1,7,10,1] 这个先序序列是非法的 )
*/
