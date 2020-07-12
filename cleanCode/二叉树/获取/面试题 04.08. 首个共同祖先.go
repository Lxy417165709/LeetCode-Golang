package 获取

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	return getLCA(root, p, q)
}

// 获取最近公共祖先，不存在则返回nil
func getLCA(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	leftSearchResult := getLCA(root.Left, p, q)
	rightSearchResult := getLCA(root.Right, p, q)

	if rightSearchResult != nil && leftSearchResult != nil {
		return root
	}
	if rightSearchResult != nil {
		return rightSearchResult
	}
	if leftSearchResult != nil {
		return leftSearchResult
	}
	return nil
}
/*
	题目链接: https://leetcode-cn.com/problems/first-common-ancestor-lcci/
	总结
		1. 在一棵树上，假如只有A，而无B，则规定最近公共祖先为A。反之亦然。 （这个规定可以减少搜索操作，即我们可以不关心B是否在该树存在）
			(这个题就是这样规定的)
*/
