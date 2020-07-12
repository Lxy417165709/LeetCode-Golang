package 性质判定

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isTwoTreeSymmetric(root.Left, root.Right)
}

func isTwoTreeSymmetric(treeA *TreeNode, treeB *TreeNode) bool {
	if treeA == nil && treeB == nil {
		return true
	}
	if treeA == nil || treeB == nil {
		return false
	}
	return treeA.Val == treeB.Val && isTwoTreeSymmetric(treeA.Left, treeB.Right) && isTwoTreeSymmetric(treeA.Right, treeB.Left)
}

/*
	题目链接: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
*/
