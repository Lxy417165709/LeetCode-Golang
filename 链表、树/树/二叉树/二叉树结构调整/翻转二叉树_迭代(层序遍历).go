package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		top := queue[0]
		top.Left, top.Right = top.Right, top.Left
		queue = queue[1:]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/invert-binary-tree/submissions/		翻转二叉树
*/
