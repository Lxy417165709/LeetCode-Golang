package main

/*
	给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
	你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
	否则不为 NULL 的节点将直接作为新二叉树的节点。
*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return mergeTreesExec(t1, t2)
}

func mergeTreesExec(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// 把t1改为t2，之后返回t2也可以的。
	t1.Left = mergeTreesExec(t1.Left, t2.Left)
	t1.Right = mergeTreesExec(t1.Right, t2.Right)
	t1.Val = t1.Val + t2.Val
	return t1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/merge-two-binary-trees/		合并二叉树
*/
