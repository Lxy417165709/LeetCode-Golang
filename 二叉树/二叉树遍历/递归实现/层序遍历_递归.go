package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var levelOrderSequence [][]int

// 递归 + 外部变量 实现层序遍历
func levelOrder(root *TreeNode) [][]int {
	levelOrderSequence = make([][]int, 0)
	levelOrderExec(root, 0)
	return levelOrderSequence
}

// lay表示当前树节点所在的层 (从0开始)
func levelOrderExec(root *TreeNode, lay int) {
	if root == nil {
		return
	}
	if len(levelOrderSequence) == lay {
		levelOrderSequence = append(levelOrderSequence, []int{})
	}
	levelOrderSequence[lay] = append(levelOrderSequence[lay], root.Val)
	levelOrderExec(root.Left, lay+1)
	levelOrderExec(root.Right, lay+1)
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/submissions/
*/
