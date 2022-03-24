package 二叉树

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror 是否互为镜像。
func isMirror(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	return A.Val == B.Val && isMirror(A.Left, B.Right) && isMirror(A.Right, B.Left)
}
