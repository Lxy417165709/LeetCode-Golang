package 二叉树

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getIsUnivalTree(root, root.Val)
}

func getIsUnivalTree(root *TreeNode, unival int) bool {
	if root == nil {
		return true
	}
	return root.Val == unival && getIsUnivalTree(root.Left, root.Val) && getIsUnivalTree(root.Right, root.Val)
}

/*
	题目链接: https://leetcode-cn.com/problems/univalued-binary-tree/
*/
