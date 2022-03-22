package 树

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSubTree(A, B) || (isSubStructure(A.Left, B)) || (isSubStructure(A.Right, B))
}

func isSubTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isSubTree(A.Left, B.Left) && isSubTree(A.Right, B.Right)
}
