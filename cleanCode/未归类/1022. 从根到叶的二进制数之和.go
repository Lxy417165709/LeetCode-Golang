package 未归类

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const mod = 1000000007

func sumRootToLeaf(root *TreeNode) int {
	return getSum(root, 0)
}

func getSum(root *TreeNode, nowSum int) int {
	if root == nil {
		return 0
	}
	nowSum = (nowSum*2 + root.Val) % mod
	if isLeaf(root) {
		return nowSum
	}
	return getSum(root.Left, nowSum) + getSum(root.Right, nowSum)
}

func isLeaf(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}

/*
	题目链接: https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/submissions/
*/
