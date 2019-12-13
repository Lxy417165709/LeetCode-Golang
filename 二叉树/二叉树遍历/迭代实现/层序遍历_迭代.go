package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代 + 队列 实现层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelOrderSequence := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		nowOrderSequence := make([]int, 0)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			nowOrderSequence = append(nowOrderSequence, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		levelOrderSequence = append(levelOrderSequence, nowOrderSequence)
	}
	return levelOrderSequence
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
*/
