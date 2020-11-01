package 结构调整

func flatten(root *TreeNode) {
	toList(root)
}

// 返回值是链表的头、尾节点 (当树空时，返回 nil, nil)
func toList(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	leftListHead, leftListTail := toList(root.Left)
	rightListHead, rightListTail := toList(root.Right)
	root.Left = nil
	switch {
	case leftListHead == nil && rightListHead == nil:
		return root, root
	case leftListHead == nil:
		root.Right = rightListHead
		return root, rightListTail
	case rightListHead == nil:
		root.Right = leftListHead
		return root, leftListTail
	default:
		root.Right = leftListHead
		leftListTail.Right = rightListHead
		return root, rightListTail
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
	总结:
		1. 之前也做过这题，这次的递归函数定义和先前的不一样。
		2. 这题最佳解法应该是迭代。
*/
