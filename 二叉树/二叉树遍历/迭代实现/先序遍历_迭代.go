package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代 + 栈 实现先序遍历
func preorderTraversal(root *TreeNode) []int {
	preorderSequence := make([]int, 0)
	isVisited := make(map[*TreeNode]int)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	isVisited[root] = 0
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			continue
		}
		if isVisited[top] == 0 {
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
			stack = append(stack, top)
			isVisited[top.Right] = 0
			isVisited[top.Left] = 0
			isVisited[top] = 1
		} else {
			preorderSequence = append(preorderSequence, top.Val)
		}
	}
	return preorderSequence
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
*/
