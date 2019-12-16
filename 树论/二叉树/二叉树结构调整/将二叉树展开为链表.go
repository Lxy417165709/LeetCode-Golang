package main

// 反后序遍历 + 外部变量 将二叉树展开为链表
var headNode *TreeNode // 链表的头节点
func flatten(root *TreeNode) {
	headNode = nil
	flattenExec(root)
}

func flattenExec(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right = flattenExec(root.Right)
	root.Left = flattenExec(root.Left)

	root.Right = headNode
	root.Left = nil
	headNode = root
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/submissions/		二叉树展开为链表
*/

/*
	总结
	1. 逆向思维很妙！ (这解法我是看别人的..)
*/
