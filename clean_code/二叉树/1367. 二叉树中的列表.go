package 二叉树

// -------------------- 错误写法 --------------------
func isSubPath(head *ListNode, anyRoot *TreeNode) bool {
	if head == nil {
		return true
	}
	if anyRoot == nil {
		return false
	}
	if head.Val == anyRoot.Val {
		return isSubPath(head.Next, anyRoot.Left) || isSubPath(head.Next, anyRoot.Right)
	} else {
		return isSubPath(head, anyRoot.Left) || isSubPath(head, anyRoot.Right)
	}
}

// -------------------- 正确写法 --------------------
func isSubPath(head *ListNode, root *TreeNode) bool {
	return isSubPathFromAnyRoot(head, root)
}

// 所谓 anyRoot，就是树里面任意一个根节点。
func isSubPathFromAnyRoot(head *ListNode, anyRoot *TreeNode) bool {
	if head == nil {
		return true
	}
	if anyRoot == nil {
		return false
	}
	return isSubPathFromCurrentRoot(head, anyRoot) || isSubPathFromAnyRoot(head, anyRoot.Left) || isSubPathFromAnyRoot(head, anyRoot.Right)
}

// 所谓 currentRoot，就是当前遍历到的根节点。
func isSubPathFromCurrentRoot(head *ListNode, currentRoot *TreeNode) bool {
	if head == nil {
		return true
	}
	if currentRoot == nil {
		return false
	}
	return head.Val == currentRoot.Val &&
		(isSubPathFromCurrentRoot(head.Next, currentRoot.Left) || isSubPathFromCurrentRoot(head.Next, currentRoot.Right))
}

/*
	题目链接: https://translate.google.cn/#view=home&op=translate&sl=zh-CN&tl=en&text=%E4%BB%BB%E6%84%8F
	总结
		1. 这题的难点在于：列表在二叉树上连续才返回true，连续很关键。
		2. 第一种做法是错误的，因为它求的是 序列在二叉树上是否连续。而第二种求的是 串在二叉树上是否连续，这就符合题意了。
		3. 这题和 _面试题 04.10. 检查子树_ 很像
*/
