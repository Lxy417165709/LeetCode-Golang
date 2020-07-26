package 路径和问题

// ---------------------- 记忆化搜索 ----------------------
const INF = 100000000000

var maxSumStartWithNode map[*TreeNode]int

func maxPathSum(root *TreeNode) int {
	maxSumStartWithNode = make(map[*TreeNode]int)
	return getMaxPathSum(root)
}

func getMaxPathSum(root *TreeNode) int {
	if root == nil {
		return -INF
	}
	maxSum := root.Val
	if leftTreeSum := maxSumStartWith(root.Left); leftTreeSum > 0 {
		maxSum += leftTreeSum
	}
	if rightTreeSum := maxSumStartWith(root.Right); rightTreeSum > 0 {
		maxSum += rightTreeSum
	}
	return max(
		maxSum,
		getMaxPathSum(root.Left),
		getMaxPathSum(root.Right),
	)
}

func maxSumStartWith(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if sum, ok := maxSumStartWithNode[root]; ok {
		return sum
	}
	maxSumStartWithNode[root] = max(
		maxSumStartWith(root.Left),
		maxSumStartWith(root.Right),
		0,
	) + root.Val

	return maxSumStartWithNode[root]
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}

// ------------- 在求取 maxSumStartWith 过程中得出答案  -------------
const INF = 100000000000

var maxSumOfPath int

func maxPathSum(root *TreeNode) int {
	maxSumOfPath = -INF
	getMaxSumStartWithAndFormMaxSumOfPath(root)
	return maxSumOfPath
}

func getMaxSumStartWithAndFormMaxSumOfPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumOfLeftTreePath := getMaxSumStartWithAndFormMaxSumOfPath(root.Left)
	sumOfRightTreePath := getMaxSumStartWithAndFormMaxSumOfPath(root.Right)
	maxSumOfPath = max(
		maxSumOfPath,
		root.Val+max(sumOfLeftTreePath, 0)+max(sumOfRightTreePath, 0),
	)
	return max(sumOfLeftTreePath, sumOfRightTreePath, 0) + root.Val
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
*/
