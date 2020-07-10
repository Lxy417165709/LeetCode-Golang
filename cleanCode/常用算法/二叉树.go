package main

// -------------------------- 获取 ---------------------------
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// -------------------------- 判断 ---------------------------
// 平衡二叉树的判断
const unBalancedFlag = -1

func isBalanced(root *TreeNode) bool {
	return getBalancedHeight(root) != unBalancedFlag
}

func getBalancedHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftBalancedHeight := getBalancedHeight(root.Left)
	rightBalancedHeight := getBalancedHeight(root.Right)

	if leftBalancedHeight == unBalancedFlag || rightBalancedHeight == unBalancedFlag {
		return unBalancedFlag
	}
	if abs(leftBalancedHeight-rightBalancedHeight) <= 1 {
		return max(rightBalancedHeight, leftBalancedHeight) + 1
	}
	return unBalancedFlag
}

// 对称二叉树的判断
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

// -------------------------- 调整内部结构 ------------------------
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
