package main

// To judge whether tree2 is the subtree of tree1
func isSubtree(tree1 *TreeNode, tree2 *TreeNode) bool {
	if isBothEmptyTree(tree1, tree2) {
		return true
	}
	if isEitherEmptyTree(tree1, tree2) {
		return false
	}
	return isSameTree(tree1, tree2) || isSubtree(tree1.Left, tree2) || isSubtree(tree1.Right, tree2)

}

// To judge whether tree1 is same to tree2
func isSameTree(tree1 *TreeNode, tree2 *TreeNode) bool {
	if isBothEmptyTree(tree1, tree2) {
		return true
	}
	if isEitherEmptyTree(tree1, tree2) {
		return false
	}
	return isSameRootVal(tree1, tree2) && isSameTree(tree1.Left, tree2.Left) && isSameTree(tree1.Right, tree2.Right)
}

func isBothEmptyTree(tree1 *TreeNode, tree2 *TreeNode) bool {
	return tree1 == nil && tree2 == nil
}
func isEitherEmptyTree(tree1 *TreeNode, tree2 *TreeNode) bool {
	return tree1 == nil || tree2 == nil
}
func isSameRootVal(tree1 *TreeNode, tree2 *TreeNode) bool {
	return tree1.Val == tree2.Val
}
