package 二叉树

// lowestCommonAncestor p，q 在树 root 中的最近公共祖先。
// 以下定义有利于递归:
// 如果 p 在 root 下，但 q 不在 root 下，此时定义最近公共祖先为 p。
// 如果 q 在 root 下，但 p 不在 root 下，此时定义最近公共祖先为 q。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 1. 递归出口。
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	// 2. 获取二者在左右子树的最近公共祖先。
	ancestorInLeftTree := lowestCommonAncestor(root.Left, p, q)
	ancestorInRightTree := lowestCommonAncestor(root.Right, p, q)

	// 3. 如果最近公共祖先分别在左右子树，此时返回根节点。 (和函数注释的定义相关)
	if ancestorInLeftTree != nil && ancestorInRightTree != nil {
		return root
	}

	// 4. 如果最近公共祖先在同在一棵树，此时返回。
	if ancestorInLeftTree != nil {
		return ancestorInLeftTree
	}
	if ancestorInRightTree != nil {
		return ancestorInRightTree
	}

	// 5. 返回。
	return nil
}
