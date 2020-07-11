package 二叉树

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val < L:
		return trimBST(root.Right, L, R)
	case root.Val > R:
		return trimBST(root.Left, L, R)
	default:
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/trim-a-binary-search-tree/
*/
