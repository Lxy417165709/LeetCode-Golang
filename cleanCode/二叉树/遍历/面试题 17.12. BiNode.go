package 遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// -------------------- 没用哑头节点 --------------------
var newHead *TreeNode
var curHead *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	newHead,curHead = nil, nil
	formNewRoot(root)
	return newHead
}

func formNewRoot(root *TreeNode) {
	if root == nil {
		return
	}
	formNewRoot(root.Left)
	if newHead== nil {
		newHead = root
		curHead = root
	} else {
		curHead.Right = root
		curHead  = curHead.Right
	}
	root.Left = nil
	formNewRoot(root.Right)
}

// -------------------------- 利用哑头节点，简化代码 --------------------------
var dummyHead *TreeNode
var curHead *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	dummyHead,curHead = &TreeNode{},nil
	curHead = dummyHead
	formDummyHead(root)
	return dummyHead.Right
}

func formDummyHead(root *TreeNode){
	if root==nil{
		return
	}
	formDummyHead(root.Left)
	curHead.Right=root
	curHead = curHead.Right
	root.Left=nil
	formDummyHead(root.Right)
}

/*
	题目链接: https://leetcode-cn.com/problems/binode-lcci/
	总结:
		1. 这题我采用了中序遍历，其实前序遍历也可以AC这题。
*/
