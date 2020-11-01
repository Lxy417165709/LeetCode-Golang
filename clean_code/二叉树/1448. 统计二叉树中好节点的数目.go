package 二叉树

const INF = 100000000000

func goodNodes(root *TreeNode) int {
	return getCountOfGoodNode(root, -INF)
}

func getCountOfGoodNode(root *TreeNode, preMaxValue int) int {
	if root == nil {
		return 0
	}
	if root.Val >= preMaxValue {
		return 1 + getCountOfGoodNode(root.Left, root.Val) + getCountOfGoodNode(root.Right, root.Val)
	} else {
		return getCountOfGoodNode(root.Left, preMaxValue) + getCountOfGoodNode(root.Right, preMaxValue)
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/
*/
