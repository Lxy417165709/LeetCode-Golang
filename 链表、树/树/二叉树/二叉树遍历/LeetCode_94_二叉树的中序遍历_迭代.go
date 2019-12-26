package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代 + 栈 实现中序遍历
func inorderTraversal(root *TreeNode) []int {
	inorderSequence := make([]int, 0)
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
			stack = append(stack, top) // 与先序遍历迭代相比，也只是调整了这句语句的位置
			stack = append(stack, top.Left)
			isVisited[top.Right] = 0
			isVisited[top] = 1
			isVisited[top.Left] = 0
		} else {
			inorderSequence = append(inorderSequence, top.Val)
		}
	}
	return inorderSequence
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/
*/
