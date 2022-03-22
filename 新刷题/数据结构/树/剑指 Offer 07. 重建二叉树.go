package 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 1. 空返回。
	if len(preorder) == 0 {
		return nil
	}

	// 2. 获取根节点在中序遍历序列的 index。
	rootNum := preorder[0]
	rootIndexInInorderSeq := 0
	for index, num := range inorder {
		if num == rootNum {
			rootIndexInInorderSeq = index
			break
		}
	}

	// 3. 递归构建。
	return &TreeNode{
		Val:   rootNum,
		Left:  buildTree(preorder[1:1+rootIndexInInorderSeq], inorder[:rootIndexInInorderSeq]),
		Right: buildTree(preorder[1+rootIndexInInorderSeq:], inorder[1+rootIndexInInorderSeq:]),
	}
}
