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

// 判断一棵树是否为完全二叉树
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	nullAppearFlag := false // 这是一个标识位，标识遍历过程中是否出现了空节点
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			nullAppearFlag = true
			continue
		}
		if nullAppearFlag {
			return false
		}
		queue = append(queue, top.Left)
		queue = append(queue, top.Right)
	}
	return true
}

/*
	题目链接:
		https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/		二叉树的完全性检验
*/

/*
	总结
	1. 由于左右子树是完全二叉树无法推出该树是完全二叉树(但完全二叉树的左右子树一定是完全二叉树)，
       所以我还没想出递归的解决方案...
	2. 以上判断采用了层序遍历的方式，判断是否出现空节点。
	   如果是完全二叉树，那么空节点后一定都是空节点，
       如果不是，那么空节点后一定会出现非空节点。
*/
