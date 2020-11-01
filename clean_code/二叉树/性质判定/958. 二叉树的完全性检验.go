package 性质判定

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftTreeHeight, rightTreeHeight := getHeight(root.Left), getHeight(root.Right)
	return leftTreeHeight == rightTreeHeight && isFullTree(root.Left) && isCompleteTree(root.Right) ||
		leftTreeHeight == rightTreeHeight+1 && isCompleteTree(root.Left) && isFullTree(root.Right)
}

func isFullTree(root *TreeNode) bool {
	countOfNode := getCountOfNode(root)
	height := getHeight(root)
	return (1<<byte(height))-1 == countOfNode
}

func getCountOfNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getCountOfNode(root.Left) + getCountOfNode(root.Right) + 1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/
	总结:
		1. 上面的代码存在大量重复操作！
		2. 可以把 getHeight、getCountOfNode 合并为一个函数。
		3. 可以通过记忆化，将树的height、countOfNode记录下来。
		4. 可以用迭代的层序遍历判断是否是完全二叉树。
*/
