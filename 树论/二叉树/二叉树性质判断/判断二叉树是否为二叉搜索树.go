package main


const (
	inf = 1000000000000
)
var lastVal int // 当前中序遍历序列的前一个数

// 判断二叉树是否为二叉搜索树
func isValidBST(root *TreeNode) bool {
	lastVal = -inf
	return isValidBSTExec(root)
}

func isValidBSTExec(root *TreeNode) bool {
	if root == nil{
		return true
	}
	if !isValidBSTExec(root.Left) || root.Val <= lastVal {
		return false
	}
	lastVal = root.Val
	return isValidBSTExec(root.Right)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/	验证二叉搜索树
*/

/*
	总结
	1. 这题本质就是中序遍历。
	2. 小贴士: 二叉搜索树的中序遍历序列是递增的。
*/