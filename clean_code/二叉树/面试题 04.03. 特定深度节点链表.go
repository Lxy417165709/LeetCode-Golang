package 二叉树

type ListNode struct {
	Val  int
	Next *ListNode
}
// ----------------------- DFS ------------------------
var listHeads []*ListNode
var listTails []*ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	listHeads = make([]*ListNode, 0)
	listTails = make([]*ListNode, 0)
	formLists(tree, 0)
	return listHeads
}

func formLists(root *TreeNode, nowDepth int) {
	if root == nil {
		return
	}
	node := &ListNode{Val: root.Val}
	if nowDepth == len(listTails) {
		listTails = append(listTails, node)
		listHeads = append(listHeads, node)
	} else {
		tail := listTails[nowDepth]
		tail.Next = node
		listTails[nowDepth] = node
	}
	formLists(root.Left, nowDepth+1)
	formLists(root.Right, nowDepth+1)
}

/*
	题目链接: https://leetcode-cn.com/problems/list-of-depth-lcci/
*/
