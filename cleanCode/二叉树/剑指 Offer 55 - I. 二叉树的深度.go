package 二叉树

func maxDepth(root *TreeNode) int {
	return getHeight(root)
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

/*
	题目链接: https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
*/
