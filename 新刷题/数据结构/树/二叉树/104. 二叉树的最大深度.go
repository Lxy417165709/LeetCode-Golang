package 二叉树

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

// maxDepth 获取二叉树深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return math_util.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
