package 性质判定

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// ------------------ 独立写的代码 ------------------
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return getIsSubStructure(A, B)
}

func getIsSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isSubStructureWithSpecificRoot(A.Left, B.Left) && isSubStructureWithSpecificRoot(A.Right, B.Right) ||
		getIsSubStructure(A.Left, B) ||
		getIsSubStructure(A.Right, B)
}

func isSubStructureWithSpecificRoot(specificRoot, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if specificRoot == nil {
		return false
	}
	return specificRoot.Val == B.Val && isSubStructureWithSpecificRoot(specificRoot.Left, B.Left) && isSubStructureWithSpecificRoot(specificRoot.Right, B.Right)
}

// ------------------ 看了官方题解后写的代码 ------------------
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return A.Val == B.Val && isSubStructureWithSpecificRoot(A.Left, B.Left) && isSubStructureWithSpecificRoot(A.Right, B.Right) ||
		isSubStructure(A.Left, B) ||
		isSubStructure(A.Right, B)
}

func isSubStructureWithSpecificRoot(specificRoot, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if specificRoot == nil {
		return false
	}
	return specificRoot.Val == B.Val && isSubStructureWithSpecificRoot(specificRoot.Left, B.Left) && isSubStructureWithSpecificRoot(specificRoot.Right, B.Right)
}

/*
	总结:
		1. 这题比较恶心的是: (约定空树不是任意一个树的子结构)
*/
