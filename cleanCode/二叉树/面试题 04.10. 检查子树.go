package 二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	return isSubTreeFromAnyRoot(t1, t2)
}

func isSubTreeFromAnyRoot(anyRoot *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if anyRoot == nil {
		return false
	}
	return isSubTreeFromCurrentRoot(anyRoot, t2) ||
		isSubTreeFromAnyRoot(anyRoot.Left, t2) ||
		isSubTreeFromAnyRoot(anyRoot.Right, t2)
}

func isSubTreeFromCurrentRoot(currentRoot *TreeNode, t2 *TreeNode) bool {
	if t2 == nil && currentRoot == nil {
		return true
	}
	if t2 == nil || currentRoot == nil {
		return false
	}
	return currentRoot.Val == t2.Val &&
		isSubTreeFromCurrentRoot(currentRoot.Left, t2.Left) &&
		isSubTreeFromCurrentRoot(currentRoot.Right, t2.Right)
}

/*
	题目链接:  https://leetcode-cn.com/problems/check-subtree-lcci/submissions/
	总结:
		1. 这题和 _1367. 二叉树中的列表_ 不一样的地方在于: 子树的叶子节点要和被查找树的叶子节点重合，
			而 _1367. 二叉树中的列表_ 不要求子链表的末节点和叶子节点重合。
		2. 这题和 _572. 另一个树的子树_ 是一样的。

*/
