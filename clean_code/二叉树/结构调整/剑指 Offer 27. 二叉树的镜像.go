package 结构调整

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root
}

/*
	题目链接: https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/submissions/
*/
