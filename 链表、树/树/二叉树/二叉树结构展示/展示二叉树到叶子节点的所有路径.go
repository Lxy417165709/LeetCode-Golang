package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 外部变量 + 递归 求二叉树所有到叶子节点的路径
var paths []string
func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	if root == nil {
		return paths
	}
	binaryTreePathsExec(root, "")
	return paths
}

func binaryTreePathsExec(root *TreeNode, nowPath string) {
	if root == nil {
		return
	}
	if nowPath != "" {
		nowPath += "->" + strconv.Itoa(root.Val)
	} else {
		nowPath += strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		paths = append(paths, nowPath)
		return
	}
	binaryTreePathsExec(root.Left, nowPath)
	binaryTreePathsExec(root.Right, nowPath)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/binary-tree-paths/		二叉树的所有路径
*/
