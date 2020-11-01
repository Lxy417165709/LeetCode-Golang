package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// ------------------------- 递归方法1 -------------------------
var sumOfAllValueOfRootToLeaf int

func sumNumbers(root *TreeNode) int {
	sumOfAllValueOfRootToLeaf = 0
	formSumOfAllValueOfRootToLeaf(root, 0)
	return sumOfAllValueOfRootToLeaf
}

func formSumOfAllValueOfRootToLeaf(root *TreeNode, curSum int) {
	if root == nil {
		return
	}
	curSum = curSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sumOfAllValueOfRootToLeaf += curSum
	}
	formSumOfAllValueOfRootToLeaf(root.Left, curSum)
	formSumOfAllValueOfRootToLeaf(root.Right, curSum)
}

// ------------------------- 递归方法2 -------------------------
func sumNumbers(root *TreeNode) int {
	return getSumOfAllValueOfRootToLeaf(root,0)
}

func getSumOfAllValueOfRootToLeaf(root *TreeNode,curSum int) int{
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val+curSum
	}
	return getSumOfAllValueOfRootToLeaf(root.Left,(curSum+root.Val)*10) + getSumOfAllValueOfRootToLeaf(root.Right,(curSum+root.Val)*10)
}
