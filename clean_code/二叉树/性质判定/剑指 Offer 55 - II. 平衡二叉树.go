package 性质判定

// --------------------------- 方法1: 时间复杂度 O(n * log_n) ---------------------------
// 概述: 在求取高度的时候，不进行平衡二叉树判断。 (与 方法2 进行对比)

const maxHeightDiffInBalancedTree = 1

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && isHeightAbsDiffOfSonTreesInBalancedRange(root)
}

func isHeightAbsDiffOfSonTreesInBalancedRange(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getSonHeightAbsDiff(root) <= maxHeightDiffInBalancedTree
}

func getSonHeightAbsDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(getHeight(root.Left) - getHeight(root.Right))
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// --------------------------- 方法2: 时间复杂度 O(n) ---------------------------
// 方法2: 在获取高度的同时，判断是否是平衡二叉树
const unBalancedFlag = -1

func isBalanced(root *TreeNode) bool {
	return getBalancedHeight(root) != unBalancedFlag
}

func getBalancedHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftBalancedHeight := getBalancedHeight(root.Left)
	rightBalancedHeight := getBalancedHeight(root.Right)

	if leftBalancedHeight == unBalancedFlag || rightBalancedHeight == unBalancedFlag {
		return unBalancedFlag
	}
	if abs(leftBalancedHeight-rightBalancedHeight) <= 1 {
		return max(rightBalancedHeight, leftBalancedHeight) + 1
	}
	return unBalancedFlag
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接: https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
	总结:
		1. 有时遍历时返回额外的信息，能够优化时间复杂度，但可能会破坏函数的单一职责原则。
		2. 我其实觉得方法1的时间复杂度是O(n^2)，但是官方给出的是 O(n * log_n)
*/

