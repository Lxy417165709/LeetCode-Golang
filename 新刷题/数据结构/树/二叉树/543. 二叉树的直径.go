package 二叉树

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := getHeightAndDiameterOfBinaryTree(root)
	return diameter
}

func getHeightAndDiameterOfBinaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, -1
	}
	leftHeight, leftDiameter := getHeightAndDiameterOfBinaryTree(root.Left)
	rightHeight, rightDiameter := getHeightAndDiameterOfBinaryTree(root.Right)
	return math_util.Max(leftHeight, rightHeight) + 1, math_util.Max(leftHeight+rightHeight, leftDiameter, rightDiameter)
}
