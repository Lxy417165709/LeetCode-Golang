package 二叉树

import (
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

const (
	FlagForNoBalanced = -1 // 不平衡标志。
)

// isBalanced 判断二叉树是否平衡。
// 平衡定义: 所有左右子树高度小于等于1。
func isBalanced(root *TreeNode) bool {
	return getSpecialHeight(root) != FlagForNoBalanced
}

// getSpecialHeight 获取特殊高度。 (当树不平衡时，返回 FlagForNoBalanced )
func getSpecialHeight(root *TreeNode) int {
	// 1. 空返回。
	if root == nil {
		return 0
	}

	// 2. 递归。
	leftSpacialHeight := getSpecialHeight(root.Left)
	rightSpacialHeight := getSpecialHeight(root.Right)

	// 3. 不平衡判断。
	if leftSpacialHeight == FlagForNoBalanced || rightSpacialHeight == FlagForNoBalanced {
		return FlagForNoBalanced
	}
	if leftSpacialHeight > rightSpacialHeight+1 || rightSpacialHeight > leftSpacialHeight+1 {
		return FlagForNoBalanced
	}

	// 4. 返回正常高度。
	return math_util.Max(leftSpacialHeight, rightSpacialHeight) + 1
}
