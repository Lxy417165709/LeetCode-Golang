package 二叉树

// ----------------------- 方法1: 求取最长直径 和 求取高度 分离 -----------------------
// 缺点: 存在重复求取相同节点高度的操作。
// 优点: 逻辑清晰。
func diameterOfBinaryTree(root *TreeNode) int {
	return getMaxDiameter(root)
}

func getMaxDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		getHeight(root.Left)+getHeight(root.Right),
		getMaxDiameter(root.Left),
		getMaxDiameter(root.Right),
	)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
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

// ----------------------- 方法2: 求取最长直径 和 求取高度 结合 -----------------------
// 缺点: 逻辑没有方法1清晰。
// 优点: 不存在重复求取相同节点高度的操作。
func diameterOfBinaryTree(root *TreeNode) int {
	_, maxDiameter := getHeightAndMaxDiameter(root)
	return maxDiameter
}
func getHeightAndMaxDiameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftTreeHeight, leftMaxDiameter := getHeightAndMaxDiameter(root.Left)
	rightTreeHeight, rightMaxDiameter := getHeightAndMaxDiameter(root.Right)
	height := max(leftTreeHeight, rightTreeHeight) + 1
	maxDiameter := max(leftMaxDiameter, rightMaxDiameter, leftTreeHeight+rightTreeHeight)
	return height, maxDiameter
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
	题目链接: https://leetcode-cn.com/problems/diameter-of-binary-tree/submissions/
	总结:
		1. 做这些题要考虑清楚: 树的最长直径是否只存在于根节点，而不在于树的子节点。
		2. 上面的做法将求取最长直径和求取高度分开了，其实在求取高度的时候就能把最长直径求出来了，
			所以可以利用这个优点进行优化。
*/
