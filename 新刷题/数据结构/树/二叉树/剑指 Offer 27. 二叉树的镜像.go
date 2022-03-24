package 二叉树

func mirrorTree(root *TreeNode) *TreeNode {
	// 1. 空返回。
	if root == nil {
		return nil
	}

	// 2. 交换左右子树。
	root.Left, root.Right = root.Right, root.Left

	// 3. 递归。
	root.Left = mirrorTree(root.Left)
	root.Right = mirrorTree(root.Right)

	// 4. 返回。
	return root
}
