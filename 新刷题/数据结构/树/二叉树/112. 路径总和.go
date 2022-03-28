package 二叉树

// hasPathSum 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true；否则，返回 false。
// 树是空时，不存在根节点到叶子节点的路径，此时返回 false。
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 1. 空树。
	if root == nil {
		return false
	}

	// 2. 叶子节点。
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// 3. 非叶子节点。
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
