package 结构调整

// ----------------- 先标记，再删除 -----------------
const removeFlag = 1000000000000

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	markRemovedLeafNodes(root, target)
	return removeMarkNode(root)
}

func removeMarkNode(root *TreeNode) *TreeNode {
	if root == nil || root.Val == removeFlag {
		return nil
	}
	root.Left = removeMarkNode(root.Left)
	root.Right = removeMarkNode(root.Right)
	return root
}

func markRemovedLeafNodes(root *TreeNode, target int) {
	if root == nil {
		return
	}
	markRemovedLeafNodes(root.Left, target)
	markRemovedLeafNodes(root.Right, target)
	if isNeedToDelete(root, target) {
		root.Val = removeFlag
	}
}

func isNeedToDelete(root *TreeNode, target int) bool {
	return isLeaf(root) && root.Val == target
}

func isLeaf(root *TreeNode) bool {
	return (root.Left == nil || root.Left.Val == removeFlag) && (root.Right == nil || root.Right.Val == removeFlag)
}

// ----------------- 直接删除 -----------------
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if isNeedToDelete(root, target) {
		return nil
	}
	return root
}

func isNeedToDelete(root *TreeNode, target int) bool {
	return isLeaf(root) && root.Val == target
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}


/*
	题目链接: https://leetcode-cn.com/problems/delete-leaves-with-a-given-value/
*/
