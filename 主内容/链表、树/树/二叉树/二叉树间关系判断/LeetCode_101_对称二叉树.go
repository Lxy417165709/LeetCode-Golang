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
// 判断二叉树是否对称
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// 判断两棵二叉树树是否互为镜像
func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/symmetric-tree/submissions/		对称二叉树
*/

/*
	总结
	1. 通过判断两棵树是否互为镜像，我们可以判断出一棵树是否是对称二叉树。
	2. 这份代码是冗余的，因为它还存在在 "二叉树间关系" 文件夹下。
*/
