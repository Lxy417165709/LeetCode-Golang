package main

// 外部变量 + 哈希查找 + 递归
var forest []*TreeNode
var shouldDelete map[int]bool
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	forest = make([]*TreeNode, 0)
	shouldDelete = make(map[int]bool)
	for i := 0; i < len(to_delete); i++ {
		shouldDelete[to_delete[i]] = true
	}
	delNodesExec(root, true)
	return forest
}

// shouldAddFlag: 该树是否需要加入结果集
func delNodesExec(root *TreeNode, shouldAddFlag bool) *TreeNode {
	if root == nil {
		return root
	}
	deleteFlag := shouldDelete[root.Val]

	// 如果根节点在to_delete中，那么我们需要把它的左右子树加入结果集。
	//          (前提是左右子树非空，这也是为什么形成结果集那里加了一个 root != nil 的判断)。
	// 如果根节点没在to_delete中，那么我们不能把它的左右子树加入结果集，而是把该根节点加入结果集。
	root.Left = delNodesExec(root.Left, deleteFlag)
	root.Right = delNodesExec(root.Right, deleteFlag)
	if deleteFlag {
		root = nil
	}
	// 形成结果集(森林)
	if shouldAddFlag && root != nil {
		forest = append(forest, root)
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/delete-nodes-and-return-forest/submissions/		删点成林
*/

/*
	总结
	1. 外部变量其实也可以使用一个引用参数代替。
	2. 这题目就是: 删除二叉树中某些值的节点，并且将剩下的二叉树记录下来。
*/
