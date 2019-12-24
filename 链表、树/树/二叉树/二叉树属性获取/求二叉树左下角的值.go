package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxDepth int // 最深层数
var ans int      // 最深层左下角的值

// 采用DFS找到二叉树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	ans = -1
	maxDepth = -1
	findBottomLeftValueExec(root, 0)
	return ans
}

func findBottomLeftValueExec(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	// 一定先遍历左子树！！
	// 假如先遍历右子树，那么得到的是二叉树右下角的值。 测试例子: [2,1,3]
	findBottomLeftValueExec(root.Left, depth+1)

	// 表示已找到叶子节点
	if root.Left == nil && root.Right == nil {
		// 判断是否是最深的
		if depth > maxDepth {
			ans = root.Val
			maxDepth = depth
		}
	}
	findBottomLeftValueExec(root.Right, depth+1)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/find-bottom-left-tree-value/submissions/		找树左下角的值
*/
/*
	总结
	1. 这题目考察的是二叉树的遍历，这题用前、中、后、层序遍历都可以做出来。
			(注意一定要先遍历左子树，因为要求的是左下角的值，如果是右下角的值的话，那么就一定要先遍历右子树)
*/
