package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// -------------------------- 方法1: 没利用二叉搜索树的性质 --------------------------
var hasPreNodeAppeared bool

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	hasPreNodeAppeared = false
	return getSuccessor(root, p)
}

func getSuccessor(root, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if leftSearchResult := getSuccessor(root.Left, p); leftSearchResult != nil {
		return leftSearchResult
	}
	if hasPreNodeAppeared {
		return root
	}
	if root.Val == p.Val {
		hasPreNodeAppeared = true
	}
	return getSuccessor(root.Right, p)
}

// -------------------------- 方法2: 利用二叉搜索树的性质 --------------------------
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return getSuccessor(root, p)
}

func getSuccessor(root, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val {
		if leftSearchResult := getSuccessor(root.Left, p); leftSearchResult != nil {
			return leftSearchResult
		}
	}
	if root.Val > p.Val {
		return root
	}
	return getSuccessor(root.Right, p)
}

// -------------------------- 方法3: 先得出中序遍历序列、再用二分法求节点 --------------------------
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return getSuccessor(root, p)
}

func getSuccessor(root, p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}
	inorderSequence := getInorderSequence(root)
	index := getIndexOfFirstGreater(inorderSequence, p.Val)
	if index == len(inorderSequence) {
		return nil
	}
	return inorderSequence[index]
}

func getInorderSequence(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	inorderSequence := make([]*TreeNode, 0)
	inorderSequence = append(inorderSequence, getInorderSequence(root.Left)...)
	inorderSequence = append(inorderSequence, root)
	inorderSequence = append(inorderSequence, getInorderSequence(root.Right)...)
	return inorderSequence
}

func getIndexOfFirstGreater(arr []*TreeNode, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid].Val > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
