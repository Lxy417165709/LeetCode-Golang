package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func rightSideView(root *TreeNode) []int {
	result = make([]int, 0)
	solve(root, 0)
	return result
}

func solve(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(result) {
		result = append(result, root.Val)
	} else {
		result[level] = root.Val
	}
	solve(root.Left, level+1)
	solve(root.Right, level+1)
}

// 总结： 这题可以采用DFS、和BFS。获取最每层最右的节点就可以了。 上面的代码采用了DFS
